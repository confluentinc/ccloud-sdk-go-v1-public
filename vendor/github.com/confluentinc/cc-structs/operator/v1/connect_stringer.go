package v1

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"strings"
)

/*
Contains overrides for String() methods so we can print more easily
If you decorate your .proto message with "option (gogoproto.goproto_stringer) = false;",
the generator will not create String() method, and you can put custom override here
*/

func (c *Connector) String() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("opv1.Connector{Name=%s, Namespace=%s, Spec={",
		c.GetName(), c.GetNamespace()))
	if c.GetSpec() == nil {
		str.WriteString("nil")
	} else {
		str.WriteString(fmt.Sprintf("pluginType=%s, desiredState=%s, UserConfigs={", c.Spec.PluginType, c.Spec.DesiredState))
		if c.Spec.Configs == nil {
			str.WriteString("nil")
		} else {
			str.WriteString(fmt.Sprintf("templateId=%s, userConfigs=%s", c.Spec.Configs.TemplateId, c.Spec.Configs.UserConfigs))
		}
		str.WriteString("}")
	}
	str.WriteString("}, Status={")
	if c.GetStatus() == nil {
		str.WriteString("nil")
	} else {
		// maybe add validatedconnectorconfigs
		str.WriteString(fmt.Sprintf("%s", c.GetStatus().GetTaskStatuses()))
	}
	str.WriteString("}}")
	return str.String()
}

// Prints all fields. Use judiciously, since this increases our log load significantly, and json
// objects are auto-indexed by elasticsearch.
func (c *Connector) verboseString() string { return proto.CompactTextString(c) }
