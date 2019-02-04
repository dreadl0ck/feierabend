/*
 * FEIERABEND - A mite integration for software developers
 * Copyright (c) 2018 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import "flag"

var (
	flagYesterday     = flag.Bool("yesterday", false, "show yesterday")
	flagDate          = flag.String("date", "", "set a date")
	flagDir           = flag.String("dir", ".", "specify a project directory")
	flagDebug         = flag.Bool("debug", false, "toggle debug mode")
	flagListProjects  = flag.Bool("projects", false, "list all available mite projects")
	flagListUsers     = flag.Bool("users", false, "list all available mite users")
	flagListCustomers = flag.Bool("customers", false, "list all available mite customers")

	// TODO: implement
	// flagUser      = flag.String("user", "", "set a user for filtering the commits")
	// flagSince = flag.String("since", "6am", "begin of workday")
	// flagUntil = flag.String("until", "", "end of workday")
)
