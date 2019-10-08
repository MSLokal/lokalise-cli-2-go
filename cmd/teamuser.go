package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	userId int64
	role   string
)

// teamUserCmd represents the team-user command
var teamUserCmd = &cobra.Command{
	Use: "team-user",
}

// teamUserListCmd represents team-user list command
var teamUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all team users. Requires Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUsers().List(teamId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserRetrieveCmd represents team-user retrieve command
var teamUserRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a Team user object. Requires Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUsers().Retrieve(teamId, userId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserUpdateCmd represents team-user update command
var teamUserUpdateCmd = &cobra.Command{
	Use: "update",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUsers().UpdateRole(teamId, userId, lokalise.TeamUserRole(role))
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserDeleteCmd represents team-user delete command
var teamUserDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes the role of a team user. Requires Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUsers().Delete(teamId, userId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	teamUserCmd.AddCommand(teamUserListCmd, teamUserRetrieveCmd, teamUserUpdateCmd, teamUserDeleteCmd)
	rootCmd.AddCommand(teamUserCmd)

	// General
	flagTeamId(teamUserCmd)

	// Update
	flagUserId(teamUserUpdateCmd)
	teamUserUpdateCmd.Flags().StringVar(&role, "role", "", "Role of the user. Available roles are owner, admin, member (required)")
	_ = teamUserUpdateCmd.MarkFlagRequired("role")

	// Retrieve, delete
	flagUserId(teamUserRetrieveCmd)
	flagUserId(teamUserDeleteCmd)

}

func flagUserId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&translationId, "user-id", 0, "A unique identifier of the user (required)")
	_ = cmd.MarkFlagRequired("user-id")
}