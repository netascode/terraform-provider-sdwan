package helpers

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

const MaxActionAttempts int = 120 // equates to 600 seconds

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetStringList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

func WaitForActionToComplete(ctx context.Context, client *sdwan.Client, id string) error {
	for attempts := 0; ; attempts++ {
		time.Sleep(5 * time.Second)
		res, err := client.Get("/device/action/status/" + id)
		if err != nil {
			return err
		}
		if res.Get("summary.status").String() == "done" {
			var failures []string
			res.Get("data").ForEach(func(k, v gjson.Result) bool {
				if strings.Contains(v.Get("statusId").String(), "failure") {
					failures = append(failures, fmt.Sprintf("Action %s for device %s failed. Activity log: %+v", id, v.Get("deviceID").String(), v.Get("activity").String()))
				}
				return true
			})
			if len(failures) > 0 {
				return errors.New(strings.Join(failures, "\n"))
			}
			break
		} else {
			tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Waiting for action '%s' to complete.", id))
		}
		if attempts > MaxActionAttempts {
			tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Maximum number of attempts reached for action '%s'.", id))
			return errors.New(fmt.Sprintf("Maximum waiting time for action '%s' reached.", id))
		}
	}
	return nil
}
