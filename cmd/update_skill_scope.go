package cmd

import (
	"fmt"
	"os"

	"github.com/nacos-group/nacos-cli/internal/help"
	"github.com/nacos-group/nacos-cli/internal/skill"
	"github.com/spf13/cobra"
)

var skillScopeValue string

var updateSkillScopeCmd = &cobra.Command{
	Use:     "skill-scope [skillName]",
	Aliases: []string{"skill-visibility"},
	Short:   "Set skill visibility scope",
	Long:    help.SkillScope.FormatForCLI("nacos-cli"),
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if skillScopeValue == "" {
			fmt.Fprintln(os.Stderr, "Error: --scope is required for skill-scope")
			os.Exit(1)
		}

		nacosClient := mustNewNacosClient()
		skillService := skill.NewSkillService(nacosClient)

		if err := skillService.UpdateSkillScope(args[0], skillScopeValue); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to update skill scope for '%s': %v\n", args[0], err)
			os.Exit(1)
		}
		fmt.Printf("Skill scope updated successfully: %s -> %s\n", args[0], skillScopeValue)
	},
}

func init() {
	updateSkillScopeCmd.Flags().StringVar(&skillScopeValue, "scope", "", "Required. Visibility scope: PUBLIC or PRIVATE")
	rootCmd.AddCommand(updateSkillScopeCmd)
}
