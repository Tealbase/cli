package integration

// Basic imports
import (
	"os"
	"testing"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	clicmd "github.com/tealbase/cli/cmd"
	"github.com/tealbase/cli/test/mocks/tealbase"
)

type MigrationTestSuite struct {
	suite.Suite
	tempDir string
	cmd     *cobra.Command
}

// test functions
func (suite *MigrationTestSuite) TestNewMigration() {
	// run command
	migration, _, err := suite.cmd.Find([]string{"migration", "new"})
	require.NoError(suite.T(), err)
	name := gonanoid.MustGenerate(tealbase.IDAlphabet, 10)
	require.NoError(suite.T(), migration.RunE(migration, []string{name}))

	// check migrations file created
	subs, err := os.ReadDir("tealbase/migrations")
	require.NoError(suite.T(), err)
	require.Regexp(suite.T(), `[0-9]{14}_`+name+".sql", subs[0].Name())
}

// hooks
func (suite *MigrationTestSuite) SetupTest() {
	// init cli
	suite.cmd = clicmd.GetRootCmd()
	suite.tempDir = NewTempDir(Logger, TempDir)

	// init tealbase
	init, _, err := suite.cmd.Find([]string{"init"})
	require.NoError(suite.T(), err)
	require.NoError(suite.T(), init.RunE(init, []string{}))
}

func (suite *MigrationTestSuite) TeardownTest() {
	require.NoError(suite.T(), os.Chdir(TempDir))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMigrationTestSuite(t *testing.T) {
	suite.Run(t, new(MigrationTestSuite))
}
