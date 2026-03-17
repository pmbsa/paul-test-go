# Project Memory

_Auto-maintained by BulletTrain. Reviewed in PRs._

## Learned from BT-1 (2026-03-17)
- Project uses Go with table-driven test patterns (21 test cases in single TestPower function)
- Integer division/truncation behavior: negative exponents return 0 (not error) due to fraction truncation to int
- Edge case convention: power of 0 returns 1 for all bases including 0 and negative numbers
