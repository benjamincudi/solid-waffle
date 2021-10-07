#!/bin/bash
set -e

createdb -U postgres waffle
psql -q -d waffle -U postgres < /tmp/structure.sql