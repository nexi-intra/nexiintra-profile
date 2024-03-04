package magicapp

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/365admin/nexiintra-profile/cmds"
)

func RegisterCmds() {
	magicbuttonCmd := &cobra.Command{
		Use:   "magicbutton",
		Short: "Magic Buttons",
		Long:  ``,
	}

	RootCmd.AddCommand(magicbuttonCmd)
	developCmd := &cobra.Command{
		Use:   "develop",
		Short: "Developer",
		Long:  ``,
	}

	RootCmd.AddCommand(developCmd)
	readCmd := &cobra.Command{
		Use:   "read",
		Short: "Configuration Data",
		Long:  ``,
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
		Short: "Compile",
		Long:  ``,
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
	CompileWashPostCmd := &cobra.Command{
		Use:   "wash",
		Short: "Wash Profile Data",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.CompileWashPost(ctx, body, args)
		},
	}
	compileCmd.AddCommand(CompileWashPostCmd)

	RootCmd.AddCommand(compileCmd)
	publishchannelsCmd := &cobra.Command{
		Use:   "publishchannels",
		Short: "Publish Channels",
		Long:  ``,
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
	PublishchannelsSasurlPostCmd := &cobra.Command{
		Use:   "sasurl",
		Short: "Generate SAS URL",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.PublishchannelsSasurlPost(ctx, args)
		},
	}
	publishchannelsCmd.AddCommand(PublishchannelsSasurlPostCmd)

	RootCmd.AddCommand(publishchannelsCmd)
	publishdomainsCmd := &cobra.Command{
		Use:   "publishdomains",
		Short: "Publish Domains",
		Long:  ``,
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
	deploywebCmd := &cobra.Command{
		Use:   "deployweb",
		Short: "Deploy Web",
		Long:  ``,
	}
	DeploywebCreatePostCmd := &cobra.Command{
		Use:   "create",
		Short: "Create Kubernetes Deployment",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.DeploywebCreatePost(ctx, body, args)
		},
	}
	deploywebCmd.AddCommand(DeploywebCreatePostCmd)

	RootCmd.AddCommand(deploywebCmd)
}
