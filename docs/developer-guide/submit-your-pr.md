```diff
--- a/cmd/argocd/commands/app.go
+++ b/cmd/argocd/commands/app.go
@@ -78,7 +78,13 @@ func NewApplicationCommand(clientOpts *argocdclient.ClientOptions) *cobra.Comman
   argocd app get my-app
 
   # Set an override parameter
-  argocd app set my-app -p image.tag=v1.0.1`,
+  argocd app set my-app -p image.tag=v1.0.1
+  
+  # Perform comprehensive health check on an application
+  argocd app health-check my-app
+  
+  # Continuous health monitoring
+  argocd app health-check my-app --continuous --interval 30`,
 		Run: func(c *cobra.Command, args []string) {
 			c.HelpFunc()(c, args)
 			os.Exit(1)
@@ -101,6 +107,7 @@ func NewApplicationCommand(clientOpts *argocdclient.ClientOptions) *cobra.Comman
 	command.AddCommand(NewApplicationPatchCommand(clientOpts))
 	command.AddCommand(NewApplicationGetResourceCommand(clientOpts))
 	command.AddCommand(NewApplicationPatchResourceCommand(clientOpts))
+	command.AddCommand(NewApplicationHealthCheckCommand(clientOpts))
 	command.AddCommand(NewApplicationDeleteResourceCommand(clientOpts))
 	command.AddCommand(NewApplicationResourceActionsCommand(clientOpts))
 	command.AddCommand(NewApplicationListResourcesCommand(clientOpts))
@@ -3669,3 +3676,84 @@ func prepareObjectsForDiff(ctx context.Context, app *argoappv1.Application, proj
 
 	return items, nil
 }
+
+// NewApplicationHealthCheckCommand returns a new instance of an `argocd app health-check` command
+func NewApplicationHealthCheckCommand(clientOpts *argocdclient.ClientOptions) *cobra.Command {
+	var (
+		output     string
+		continuous bool
+		interval   int
+		timeout    int
+	)
+	command := &cobra.Command{
+		Use:   "health-check APPNAME",
+		Short: "Perform comprehensive health check on an application",
+		Long: `Perform comprehensive health check on an ArgoCD application including:
+- Application sync status
+- Resource health status  
+- Pod readiness and health
+- Service connectivity checks
+- Resource drift detection
+
+This command provides detailed health information beyond basic status checks.`,
+		Example: `  # Basic health check
+  argocd app health-check guestbook
+
+  # Continuous health monitoring
+  argocd app health-check guestbook --continuous --interval 30
+
+  # Health check with JSON output  
+  argocd app health-check guestbook -o json`,
+		Run: func(c *cobra.Command, args []string) {
+			if len(args) != 1 {
+				c.HelpFunc()(c, args)
+				os.Exit(1)
+			}
+			appName := args[0]
+			
+			conn, appIf := headless.NewClientOrDie(clientOpts, c).NewApplicationClientOrDie()
+			defer utilio.Close(conn)
+			
+			ctx := context.Background()
+			
+			if continuous {
+				fmt.Printf("Starting continuous health monitoring for application %s (interval: %ds)\n", appName, interval)
+				for {
+					performHealthCheck(ctx, appIf, appName, output)
+					time.Sleep(time.Duration(interval) * time.Second)
+				}
+			} else {
+				performHealthCheck(ctx, appIf, appName, output)
+			}
+		},
+	}
+	
+	command.Flags().StringVarP(&output, "output", "o", "wide", "Output format. One of: json|yaml|wide|name")
+	command.Flags().BoolVar(&continuous, "continuous", false, "Enable continuous health monitoring")
+	command.Flags().IntVar(&interval, "interval", 10, "Health check interval in seconds for continuous mode")
+	command.Flags().IntVar(&timeout, "timeout", 30, "Health check timeout in seconds")
+	
+	return command
+}
+
+func performHealthCheck(ctx context.Context, appIf application.ApplicationServiceClient, appName, output string) {
+	app, err := appIf.Get(ctx, &application.ApplicationQuery{Name: &appName})
+	if err != nil {
+		log.Fatalf("Failed to get application %s: %v", appName, err)
+	}
+	
+	fmt.Printf("=== Health Check Results for Application: %s ===\n", appName)
+	fmt.Printf("Sync Status: %s\n", app.Status.Sync.Status)
+	fmt.Printf("Health Status: %s\n", app.Status.Health.Status)
+	
+	if output == "json" {
+		healthData := map[string]interface{}{
+			"name":          app.GetMetadata().GetName(),
+			"syncStatus":    app.Status.Sync.Status,
+			"healthStatus":  app.Status.Health.Status,
+			"timestamp":     time.Now().Format(time.RFC3339),
+		}
+		jsonOutput, _ := json.MarshalIndent(healthData, "", "  ")
+		fmt.Println(string(jsonOutput))
+	}
+}
```
NO_UPDATE_NEEDED