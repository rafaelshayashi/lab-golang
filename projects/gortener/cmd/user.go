/*
Copyright © 2024 Rafael Hayashi <rafael@hayashi.dev>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lab-golang/projects/gortener/database"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var (
	userCmd = &cobra.Command{
		Use:   "user",
		Short: "user related commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("user called")
		},
	}
	addUserCmd = &cobra.Command{
		Use:   "add",
		Short: "add a new user",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add user called")

			username, _ := cmd.Flags().GetString("username")
			name, _ := cmd.Flags().GetString("name")
			email, _ := cmd.Flags().GetString("email")

			connectionURL := "postgres://postgres:changeme@localhost:5432/gortener?sslmode=disable"
			connection, err := pgx.Connect(context.Background(), connectionURL)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
				os.Exit(1)
			}
			defer connection.Close(context.Background())
			query := database.New(connection)

			id, err := query.CreateUser(context.Background(), database.CreateUserParams{
				Username: pgtype.Text{String: username, Valid: true},
				Name:     pgtype.Text{String: name, Valid: true},
				Email:    pgtype.Text{String: email, Valid: true},
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("user created id: ", id)
		},
	}
)

func init() {
	addUserCmd.Flags().String("username", "", "username")
	addUserCmd.Flags().String("name", "", "name")
	addUserCmd.Flags().String("email", "", "email")
	userCmd.AddCommand(addUserCmd)
	rootCmd.AddCommand(userCmd)
}
