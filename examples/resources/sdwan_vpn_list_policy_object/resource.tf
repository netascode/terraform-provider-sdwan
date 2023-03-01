resource "sdwan_vpn_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      vpn_id = 1
    }
  ]
}
