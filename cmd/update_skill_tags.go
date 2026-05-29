package cmd

import (
	"fmt"
	"os"

	"github.com/nacos-group/nacos-cli/internal/help"
	"github.com/nacos-group/nacos-cli/internal/skill"
	"github.com/spf13/cobra"
)

var skillTagsValue string

var updateSkillTagsCmd = &cobra.Command{
	Use:     "skill-tags [skillName]",
	Aliases: []string{"skill-biz-tags"},
	Short:   "Set skill metadata tags",
	Long:    help.SkillTags.FormatForCLI("nacos-cli"),
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if skillTagsValue == "" {
			fmt.Fprintln(os.Stderr, "Error: --tags is required for skill-tags")
			os.Exit(1)
		}

		nacosClient := mustNewNacosClient()
		skillService := skill.NewSkillService(nacosClient)

		if err := skillService.UpdateSkillBizTags(args[0], skillTagsValue); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to update skill tags for '%s': %v\n", args[0], err)
			os.Exit(1)
		}
		fmt.Printf("Skill tags updated successfully: %s -> %s\n", args[0], skillTagsValue)
	},
}

func init() {
	updateSkillTagsCmd.Flags().StringVar(&skillTagsValue, "tags", "", "Required. Skill metadata tags, for example retail,finance")
	rootCmd.AddCommand(updateSkillTagsCmd)
}
