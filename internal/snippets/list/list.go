package list

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-errors/errors"
	"github.com/spf13/afero"
	"github.com/tealbase/cli/internal/migration/list"
	"github.com/tealbase/cli/internal/utils"
	"github.com/tealbase/cli/internal/utils/flags"
	"github.com/tealbase/cli/pkg/api"
)

func Run(ctx context.Context, fsys afero.Fs) error {
	ref, err := flags.LoadProjectRef(fsys)
	if err != nil {
		return err
	}
	resp, err := utils.GetTealbase().ListSnippetsWithResponse(ctx, &api.ListSnippetsParams{ProjectRef: &ref})
	if err != nil {
		return errors.Errorf("failed to list snippets: %w", err)
	}

	if resp.JSON200 == nil {
		return errors.New("Unexpected error listing SQL snippets: " + string(resp.Body))
	}

	table := `|ID|NAME|VISIBILITY|OWNER|CREATED AT (UTC)|UPDATED AT (UTC)|
|-|-|-|-|-|-|
`
	for _, snippet := range resp.JSON200.Data {
		table += fmt.Sprintf(
			"|`%s`|`%s`|`%s`|`%s`|`%s`|`%s`|\n",
			snippet.Id,
			strings.ReplaceAll(snippet.Name, "|", "\\|"),
			strings.ReplaceAll(string(snippet.Visibility), "|", "\\|"),
			strings.ReplaceAll(snippet.Owner.Username, "|", "\\|"),
			utils.FormatTimestamp(snippet.InsertedAt),
			utils.FormatTimestamp(snippet.UpdatedAt),
		)
	}

	return list.RenderTable(table)
}
