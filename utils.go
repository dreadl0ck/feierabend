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

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// exit with a message, an error and optional values
func exitWith(message string, err error, values ...string) {
	var arr = []string{message, err.Error()}
	fmt.Println(strings.Join(append(arr, values...), " "))
	os.Exit(1)
}

// reset a time.Time to the exact beginning of the day
// e.g: 2019-02-04 00:00:00
func zeroDate(t time.Time) time.Time {
	t = t.Add(-time.Duration(t.Hour()) * time.Hour)
	t = t.Add(-time.Duration(t.Minute()) * time.Minute)
	t = t.Add(-time.Duration(t.Second()) * time.Second)
	t = t.Add(-time.Duration(t.Nanosecond()) * time.Nanosecond)
	return t
}

func debug(values ...interface{}) {
	if *flagDebug {
		fmt.Print("[DEBUG]:")
		fmt.Println(values...)
	}
}
