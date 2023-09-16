//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageprovider

import(
"fmt"
"time"

"github.com/YakDriver/regexache"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

//validAssumeRoleDurationvalidatesastringcanbeparsedasavalidtime.Duration
//andiswithinaminimumof15minutesandmaximumof12hours
funcvalidAssumeRoleDuration(vinterface{},kstring)(ws[]string,errors[]error){
duration,err:=time.ParseDuration(v.(string))

iferr!=nil{
errors=append(errors,fmt.Errorf("%qcannotbeparsedasaduration:%w",k,err))
return
}

ifduration.Minutes()<15||duration.Hours()>12{
errors=append(errors,fmt.Errorf("duration%qmustbebetween15minutes(15m)and12hours(12h),inclusive",k))
}

return
}

varvalidAssumeRoleSessionName=validation.All(
validation.StringLenBetween(2,64),
validation.StringMatch(regexache.MustCompile(`[\w+=,.@\-]*`),""),
)

varvalidAssumeRoleSourceIdentity=validation.All(
validation.StringLenBetween(2,64),
validation.StringMatch(regexache.MustCompile(`[\w+=,.@\-]*`),""),
)
