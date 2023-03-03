resource "sdwan_sla_class_policy_object" "example" {
  name = "Example"
  entries = [
    {
      app_probe_class_id            = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      jitter                        = 100
      latency                       = 10
      loss                          = 1
      fallback_best_tunnel_criteria = "jitter-loss-latency"
      fallback_best_tunnel_jitter   = 100
      fallback_best_tunnel_latency  = 10
      fallback_best_tunnel_loss     = 1
    }
  ]
}
