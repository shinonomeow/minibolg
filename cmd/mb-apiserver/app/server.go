package app

import (
	"github.com/shinonomeow/miniblog/cmd/mb-apiserver/app/options"
	"github.com/spf13/cobra"
)

var configFile string

func NewMiniBlogCommand() *cobra.Command {
	opts := options.NewServerOptions()
	cmd := &cobra.Command{
		Use:   "mb-apiserver",
		Short: "A mini blog show best parctices for develop a full-featured Go project",
		Long: `A mini blog show best practices for develop a full-featured Go project.
		The project features include:
		• Utilization of a clean architecture;
		• Use of many commonly used Go packages: gorm, casbin, govalidator, jwt, gin, 
			cobra, viper, pflag, zap, pprof, grpc, protobuf, grpc-gateway, etc.;
		• A standardized directory structure following the project-layout convention;
		• Authentication (JWT) and authorization features (casbin);
		• Independently designed log and error packages;
		• Management of the project using a high-quality Makefile;
		• Static code analysis;
		• Includes unit tests, performance tests, fuzz tests, and mock tests;
		• Rich web functionalities (tracing, graceful shutdown, middleware, CORS, 
			recovery from panics, etc.);
		• Implementation of HTTP, HTTPS, and gRPC servers;
		• Implementation of JSON and Protobuf data exchange formats;
		• The project adheres to numerous development standards: 
			code standards, versioning standards, API standards, logging standards, 
			error handling standards, commit standards, etc.;
		• Access to MySQL with programming implementation;
		• Implemented business functionalities: user management and blog management;
		• RESTful API design standards;
		• OpenAPI 3.0/Swagger 2.0 API documentation;
		• High-quality code.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
	}
	cobra.OnInitialize(onInitialize)
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the miniblog configuration file.")
	return cmd
}

func run(opts *options.ServerOptions) error {
	cfg, err := opts.Config()
	if err != nil {
		return err
	}
	server, err := cfg.NewUnionServer()
	if err != nil {
		return err
	}
	return server.Run()
}
