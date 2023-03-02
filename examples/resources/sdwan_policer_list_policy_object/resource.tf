resource "sdwan_policer_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      burst         = 100000
      exceed_action = "remark"
      rate          = 100
    }
  ]
}
