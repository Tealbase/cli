package list

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-errors/errors"
	"github.com/tealbase/cli/internal/migration/list"
	"github.com/tealbase/cli/internal/utils"
)

func Run(ctx context.Context) error {
	resp, err := utils.GetTealbase().GetOrganizationsWithResponse(ctx)
	if err != nil {
		return errors.Errorf("failed to list organizations: %w", err)
	}

	if resp.JSON200 == nil {
		return errors.New("Unexpected error retrieving organizations: " + string(resp.Body))
	}

	table := `|ID|NAME|
|-|-|
`
	for _, org := range *resp.JSON200 {
		table += fmt.Sprintf("|`%s`|`%s`|\n", org.Id, strings.ReplaceAll(org.Name, "|", "\\|"))
	}

	return list.RenderTable(table)
}
