# Licensed under the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License. You may obtain
# a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.

import io
import tempfile

import fixtures
import testtools

from burstbuffer.cmd import fakewarp


class TestFakeWarp(testtools.TestCase):

    def setUp(self):
        super(TestFakeWarp, self).setUp()
        self.stdout = io.StringIO()
        self.useFixture(fixtures.MonkeyPatch('sys.stdout', self.stdout))
        self.stderr = io.StringIO()
        self.useFixture(fixtures.MonkeyPatch('sys.stderr', self.stderr))

    def test_pools(self):
        result = fakewarp.main(["--function", "pools"])
        self.assertEqual(0, result)

    def test_show_instances(self):
        result = fakewarp.main(["--function", "show_instances"])
        self.assertEqual(0, result)

    def test_show_sessions(self):
        result = fakewarp.main(["--function", "show_sessions"])
        self.assertEqual(0, result)

    def test_teardown(self):
        cmdline = "--function teardown --token 347 --job /tmp/fakescript"
        result = fakewarp.main(cmdline.split(" "))
        self.assertEqual(0, result)

    def test_job_process(self):
        with tempfile.NamedTemporaryFile() as f:
            f.write(b"#!/bin/bash\n")
            f.write(b"#DW jobdw capacity=1GiB")
            f.flush()
            cmdline = "--function job_process --job %s" % f.name

            result = fakewarp.main(cmdline.split(" "))

            self.assertEqual(0, result)
            self.assertEqual("capacity=1GiB\n", self.stdout.getvalue())