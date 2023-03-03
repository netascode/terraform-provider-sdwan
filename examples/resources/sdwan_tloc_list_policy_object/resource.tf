resource "sdwan_tloc_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      color         = "blue"
      encapsulation = "gre"
      preference    = 10
      tloc_ip       = "1.1.1.2"
    }
  ]
}
