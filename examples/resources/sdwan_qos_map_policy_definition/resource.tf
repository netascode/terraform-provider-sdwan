resource "sdwan_qos_map_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  qos_schedulers = [
    {
      bandwidth_percent = 10
      buffer_percent    = 10
      burst             = 100000
      class_map_id      = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      drop_type         = "red-drop"
      queue             = 6
      scheduling_type   = "wrr"
    }
  ]
}
