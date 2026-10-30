#!/bin/sh
# probe-kit 시작 — 과제 폴더(README.md가 있는 곳)의 터미널에서 한 줄:
#   git clone https://github.com/Taewoo-Won/probe-kit /tmp/harness && sh /tmp/harness/go
# 하는 일은 셋뿐이다: AGENTS.md 맨 아래에 하네스 규칙 붙이기 · MSG.md 만들기 · 다음 할 일 보여 주기.
# 설치 · 네트워크 · git 커밋 · Codex 설정은 건드리지 않는다. 여러 번 쳐도 같은 결과다.
H=/tmp/harness
MARK='한 단어만 치면 아래 정의대로'

if [ ! -f "$H/AGENTS.md" ] || [ ! -f "$H/MSG.md" ]; then
  echo "[하네스] $H 에 AGENTS.md · MSG.md가 없다. clone부터: git clone https://github.com/Taewoo-Won/probe-kit $H"
  exit 1
fi

if [ "$1" != "f" ] && ! ls -1 | grep -qi '^readme'; then
  echo "[하네스] 여기($(pwd))에 README가 없어 아무것도 안 했다."
  echo "[하네스] 과제 폴더로 옮긴 뒤 다시: sh $H/go   (여기가 과제 폴더가 맞으면: sh $H/go f)"
  exit 1
fi

# 1) AGENTS.md — 과제가 준 AGENTS.md가 있으면 지우지 않고 맨 아래에 붙인다(과제 규칙이 우선)
if [ -f AGENTS.md ] && grep -q "$MARK" AGENTS.md; then
  echo "[하네스] AGENTS.md: 하네스 규칙이 이미 있어 그대로 둠"
elif [ -f AGENTS.md ]; then
  printf '\n' >> AGENTS.md
  cat "$H/AGENTS.md" >> AGENTS.md
  echo "[하네스] AGENTS.md: 과제 규칙 아래에 하네스 규칙을 붙임 (과제 규칙이 우선)"
else
  cat "$H/AGENTS.md" > AGENTS.md
  echo "[하네스] AGENTS.md: 새로 만듦"
fi

# 2) MSG.md — 내 메모. 이미 있으면 덮어쓰지 않는다
if [ -f MSG.md ] && grep -q '^# MSG.md' MSG.md; then
  echo "[하네스] MSG.md: 이미 있어 그대로 둠 (새로 받으려면: cat $H/MSG.md > MSG.md)"
elif [ -f MSG.md ]; then
  echo "[하네스] MSG.md: 과제에 같은 이름 파일이 있어 건드리지 않음 → 내 메모는 $H/MSG.md"
else
  cat "$H/MSG.md" > MSG.md
  echo "[하네스] MSG.md: 만듦 (내 메모 · 커밋 안 함)"
fi

# 3) 다음 할 일
echo "[하네스] 다음: 같은 탭에 codex → 'Update available'이면 2 · 'Do you trust'면 Enter → /model Sol 6.1 medium → /status → 정찰"
echo "[하네스] 흐름 · 문장은 편집기에서 MSG.md (맨 위 '흐름')"
