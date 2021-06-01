// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !linux

package linux

import (
	"io"

	"github.com/banzaicloud/pke/cmd/pke/app/constants"
)

func CentOSVersion(w io.Writer) (string, error) {
	return "", constants.ErrUnsupportedOS
}

func RedHatVersion(w io.Writer) (string, error) {
	return "", constants.ErrUnsupportedOS
}

func LSBReleaseDistributorID(w io.Writer) (string, error) {
	return "", constants.ErrUnsupportedOS
}

func LSBReleaseReleaseNumber(w io.Writer) (string, error) {
	return "", constants.ErrUnsupportedOS
}
