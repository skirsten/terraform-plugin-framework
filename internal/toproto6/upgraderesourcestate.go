package toproto6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// UpgradeResourceStateResponse returns the *tfprotov6.UpgradeResourceStateResponse
// equivalent of a *fwserver.UpgradeResourceStateResponse.
func UpgradeResourceStateResponse(ctx context.Context, fw *fwserver.UpgradeResourceStateResponse) *tfprotov6.UpgradeResourceStateResponse {
	if fw == nil {
		return nil
	}

	proto6 := &tfprotov6.UpgradeResourceStateResponse{
		Diagnostics: Diagnostics(fw.Diagnostics),
	}

	upgradedState, diags := State(ctx, fw.UpgradedState)

	proto6.Diagnostics = append(proto6.Diagnostics, Diagnostics(diags)...)
	proto6.UpgradedState = upgradedState

	return proto6
}