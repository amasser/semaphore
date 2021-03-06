package hcl

import (
	"math/big"
	"testing"

	"github.com/jexia/semaphore/pkg/broker"
	"github.com/jexia/semaphore/pkg/broker/logger"
	"github.com/jexia/semaphore/pkg/specs"
	"github.com/jexia/semaphore/pkg/specs/types"
	"github.com/zclconf/go-cty/cty"
)

func TestSetDefaultValue(t *testing.T) {
	type expected struct {
		Default interface{}
		Type    types.Type
	}

	tests := map[cty.Value]expected{
		cty.StringVal("default"): {
			Default: "default",
			Type:    types.String,
		},
		cty.NumberVal(big.NewFloat(10)): {
			Default: int64(10),
			Type:    types.Int64,
		},
		cty.BoolVal(true): {
			Default: true,
			Type:    types.Bool,
		},
	}

	for input, expected := range tests {
		ctx := logger.WithLogger(broker.NewBackground())
		property := specs.Property{}
		SetDefaultValue(ctx, &property, input)

		if expected.Default != property.Default {
			t.Errorf("unexpected result %+v, expected %+v", property.Default, expected.Default)
		}

		if expected.Type != property.Type {
			t.Errorf("unexpected type %s, expected %s", property.Type, expected.Type)
		}
	}
}
