package magicapp

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/365admin/nexiintra-profile/cmds"
)

func RegisterCmds() {
	readCmd := &cobra.Command{
		Use:   "read",
		Short: "20-read",
		Long:  `Describe the main purpose of this kitchen`,
	}
	ReadCountriesPostCmd := &cobra.Command{
		Use:   "countries",
		Short: "Read Countries",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ReadCountriesPost(ctx, args)
		},
	}
	readCmd.AddCommand(ReadCountriesPostCmd)
	ReadChannelsPostCmd := &cobra.Command{
		Use:   "channels",
		Short: "Read Channels",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ReadChannelsPost(ctx, args)
		},
	}
	readCmd.AddCommand(ReadChannelsPostCmd)
	ReadCategoriesPostCmd := &cobra.Command{
		Use:   "categories",
		Short: "Read Categories",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ReadCategoriesPost(ctx, args)
		},
	}
	readCmd.AddCommand(ReadCategoriesPostCmd)
	ReadUnitsPostCmd := &cobra.Command{
		Use:   "units",
		Short: "Read Units",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ReadUnitsPost(ctx, args)
		},
	}
	readCmd.AddCommand(ReadUnitsPostCmd)
	ReadWhilelisteddomainsPostCmd := &cobra.Command{
		Use:   "whilelisteddomains",
		Short: "Read Whitelisted Guest Domains",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ReadWhilelisteddomainsPost(ctx, args)
		},
	}
	readCmd.AddCommand(ReadWhilelisteddomainsPostCmd)

	RootCmd.AddCommand(readCmd)
	compileCmd := &cobra.Command{
		Use:   "compile",
		Short: "30-compile",
		Long:  `Describe the main purpose of this kitchen`,
	}
	CompileLookupvaluesPostCmd := &cobra.Command{
		Use:   "lookupvalues",
		Short: "Compile Lookup values",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.CompileLookupvaluesPost(ctx, args)
		},
	}
	compileCmd.AddCommand(CompileLookupvaluesPostCmd)

	RootCmd.AddCommand(compileCmd)
	publishchannelsCmd := &cobra.Command{
		Use:   "publishchannels",
		Short: "40-publish-channels",
		Long:  `Describe the main purpose of this kitchen`,
	}
	PublishchannelsUploadblobPostCmd := &cobra.Command{
		Use:   "uploadblob",
		Short: "Publish Blob Storage",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.PublishchannelsUploadblobPost(ctx, body, args)
		},
	}
	publishchannelsCmd.AddCommand(PublishchannelsUploadblobPostCmd)
	PublishchannelsGeturlPostCmd := &cobra.Command{
		Use:   "geturl",
		Short: "Publish Blob Storage",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.PublishchannelsGeturlPost(ctx, args)
		},
	}
	publishchannelsCmd.AddCommand(PublishchannelsGeturlPostCmd)

	RootCmd.AddCommand(publishchannelsCmd)
	publishdomainsCmd := &cobra.Command{
		Use:   "publishdomains",
		Short: "41-publish-domains",
		Long:  `Describe the main purpose of this kitchen`,
	}
	PublishdomainsUploadblobPostCmd := &cobra.Command{
		Use:   "uploadblob",
		Short: "Publish Blob Storage",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.PublishdomainsUploadblobPost(ctx, body, args)
		},
	}
	publishdomainsCmd.AddCommand(PublishdomainsUploadblobPostCmd)
	PublishdomainsGeturlPostCmd := &cobra.Command{
		Use:   "geturl",
		Short: "Get SAS URL",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.PublishdomainsGeturlPost(ctx, args)
		},
	}
	publishdomainsCmd.AddCommand(PublishdomainsGeturlPostCmd)

	RootCmd.AddCommand(publishdomainsCmd)
	developCmd := &cobra.Command{
		Use:   "develop",
		Short: "50-develop",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(developCmd)
}
