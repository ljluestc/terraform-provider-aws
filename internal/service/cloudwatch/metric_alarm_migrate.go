//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagecloudwatch

import(
"fmt"
"log"

"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

funcMetricAlarmMigrateState(
vint,is*terraform.InstanceState,metainterface{})(*terraform.InstanceState,error){
switchv{
case0:
log.Println("[INFO]FoundAWSCloudWatchMetricAlarmStatev0;migratingtov1")
returnmigrateMetricAlarmStateV0toV1(is)
default:
returnis,fmt.Errorf("Unexpectedschemaversion:%d",v)
}
}

funcmigrateMetricAlarmStateV0toV1(is*terraform.InstanceState)(*terraform.InstanceState,error){
ifis.Empty(){
log.Println("[DEBUG]EmptyInstanceState;nothingtomigrate.")
returnis,nil
}

log.Printf("[DEBUG]Attributesbeforemigration:%#v",is.Attributes)

is.Attributes["treat_missing_data"]="missing"

log.Printf("[DEBUG]Attributesaftermigration:%#v",is.Attributes)
returnis,nil
}
