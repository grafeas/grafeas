// Copyright 2019 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"fmt"

	cpb "github.com/grafeas/grafeas/proto/v1beta1/common_go_proto"
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	pkgpb "github.com/grafeas/grafeas/proto/v1beta1/package_go_proto"
	vpb "github.com/grafeas/grafeas/proto/v1beta1/vulnerability_go_proto"
)

const (
	TestNoteID = "CVE-1999-0710"
)

func TestOccurrence(pID, noteName string) *pb.Occurrence {
	return &pb.Occurrence{
		Name:     fmt.Sprintf("projects/%s/occurrences/134", pID),
		Resource: &pb.Resource{Uri: "gcr.io/foo/bar"},
		NoteName: noteName,
		Kind:     cpb.NoteKind_VULNERABILITY,
		Details: &pb.Occurrence_Vulnerability{
			Vulnerability: &vpb.Details{
				Severity:  vpb.Severity_HIGH,
				CvssScore: 7.5,
				PackageIssue: []*vpb.PackageIssue{
					{
						SeverityName: "HIGH",
						AffectedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:8",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "52.1",
								Revision: "8+deb8u3",
							},
						},
						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:8",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "52.1",
								Revision: "8+deb8u4",
							},
						},
					},
				},
			},
		},
	}
}

func TestNote(pID string) *pb.Note {
	return &pb.Note{
		Name:             fmt.Sprintf("projects/%s/notes/%s", pID, TestNoteID),
		ShortDescription: "CVE-2014-9911",
		LongDescription:  "NIST vectors: AV:N/AC:L/Au:N/C:P/I:P",
		Kind:             cpb.NoteKind_VULNERABILITY,
		Type: &pb.Note_Vulnerability{
			&vpb.Vulnerability{
				CvssScore: 7.5,
				Severity:  vpb.Severity_HIGH,
				Details: []*vpb.Vulnerability_Detail{
					{
						CpeUri:  "cpe:/o:debian:debian_linux:7",
						Package: "icu",
						Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
							"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
							"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
						MinAffectedVersion: &pkgpb.Version{
							Kind: pkgpb.Version_MINIMUM,
						},
						SeverityName: "HIGH",

						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:7",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "4.8.1.1",
								Revision: "12+deb7u6",
							},
						},
					},
					{
						CpeUri:  "cpe:/o:debian:debian_linux:8",
						Package: "icu",
						Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
							"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
							"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
						MinAffectedVersion: &pkgpb.Version{
							Kind: pkgpb.Version_MINIMUM,
						},
						SeverityName: "HIGH",

						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:8",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "52.1",
								Revision: "8+deb8u4",
							},
						},
					},
					{
						CpeUri:  "cpe:/o:debian:debian_linux:9",
						Package: "icu",
						Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
							"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
							"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
						MinAffectedVersion: &pkgpb.Version{
							Kind: pkgpb.Version_MINIMUM,
						},
						SeverityName: "HIGH",

						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:9",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "55.1",
								Revision: "3",
							},
						},
					},
					{
						CpeUri:  "cpe:/o:canonical:ubuntu_linux:14.04",
						Package: "android",
						Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
							"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
							"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
						MinAffectedVersion: &pkgpb.Version{
							Kind: pkgpb.Version_MINIMUM,
						},
						SeverityName: "MEDIUM",

						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:canonical:ubuntu_linux:14.04",
							Package: "android",
							Version: &pkgpb.Version{
								Kind: pkgpb.Version_MINIMUM,
							},
						},
					},
				},
			},
		},
		RelatedUrl: []*cpb.RelatedUrl{
			{
				Url:   "https://security-tracker.debian.org/tracker/CVE-2014-9911",
				Label: "More Info",
			},
			{
				Url:   "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9911",
				Label: "More Info",
			},
		},
	}
}
