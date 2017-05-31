# -*- coding: utf-8 -*-

import os
import os.path
import re
import sys

_UPDATE_REPORT_RE = re.compile(
    r'^.+updateReport:\s+func\(.+\)\s*{\s*\w+.(\w+)\s*=\s*\w+\s*}', re.MULTILINE)
_NAME_RE = re.compile(r'^.+Name:\s*"(\w+)"', re.MULTILINE)


def clean(var):
    return var.replace("_", "").lower()

def meta_lint(fd):
    ok = True
    assigned_to = None
    name = None
    full = fd.read()
    m = _UPDATE_REPORT_RE.search(full)
    if m:
        assigned_to = m.group(1)
    m = _NAME_RE.search(full)
    if m:
        name = m.group(1)
    if clean(assigned_to) != clean(name):
        msg = "mismatch name ({!s}) vs report assignment ({!s})\n".format(
            name, assigned_to)
        sys.stderr.write(msg)
        ok = False
    sys.stderr.flush()
    return ok


def main():
    dirname = os.path.dirname(__file__)
    ok = True
    for leaf in os.listdir(dirname):
        full_path = os.path.join(dirname, leaf)
        if leaf.endswith(".go") \
                and leaf.startswith("lint_") \
                and not leaf.endswith("_test.go") \
                and os.path.isfile(full_path):
            with open(full_path) as fd:
                if not meta_lint(fd):
                    ok = False
    if not ok:
        sys.exit(1)


if __name__ == "__main__":
    main()
