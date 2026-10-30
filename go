#!/bin/sh
# probe-kit 시작 — 과제 폴더의 터미널에서 한 줄:
#   git clone https://github.com/Taewoo-Won/probe-kit /tmp/harness; sh /tmp/harness/go
# 하는 일: AGENTS.md 맨 아래에 하네스 규칙 붙이기 · MSG.md 만들기 · 다음 할 일 보여 주기
# (작업 폴더 안에 잘못 받은 하네스 사본이 있으면 지운다). 설치 · 네트워크 · git 커밋 · Codex 설정은 건드리지 않는다.
# 여러 번 쳐도 결과가 같다. 그래서 한 줄을 통째로 다시 쳐도 된다.
H=/tmp/harness
MARK='한 단어만 치면 아래 정의대로'
say() { echo "[하네스] $*"; }

if [ ! -f "$H/AGENTS.md" ] || [ ! -f "$H/MSG.md" ]; then
  say "$H 에 하네스가 없다(clone이 안 됐다). 주소를 확인하고 한 줄을 다시: git clone https://github.com/Taewoo-Won/probe-kit $H; sh $H/go"
  exit 1
fi

# 작업 폴더 안에 잘못 받은 하네스 사본(./tmp/harness · ./probe-kit · ./harness)은 지운다 — Codex가 읽거나 제출물에 섞이지 않게
for d in tmp/harness probe-kit harness; do
  if [ -f "$d/go" ] && [ -d "$d/.git" ] && grep -q "$MARK" "$d/AGENTS.md" 2>/dev/null; then
    rm -rf "$d" && say "작업 폴더 안에 잘못 받은 $d 를 지움"
    [ "$d" = tmp/harness ] && rmdir tmp 2>/dev/null
  fi
done

# 과제 폴더인지: README · .git · AGENTS.md 중 하나라도 있으면 과제 폴더로 본다
here=no
ls -1A 2>/dev/null | grep -qi '^readme' && here=yes
[ -d .git ] && here=yes
[ -f AGENTS.md ] && here=yes
if [ "$1" != "f" ] && [ "$here" = no ]; then
  sub=""
  for x in */; do
    [ -d "$x" ] || continue
    if ls -1A "$x" 2>/dev/null | grep -qi '^readme'; then sub="${x%/}"; break; fi
  done
  say "여기($(pwd))에 README가 없어 아무것도 안 했다."
  if [ -n "$sub" ]; then
    say "과제 폴더로 옮겨서 다시: cd $sub; sh $H/go   (그다음 codex도 같은 탭에서)"
  else
    say "과제 폴더(README가 있는 곳)로 cd 한 뒤 다시: sh $H/go   (여기가 과제 폴더가 맞으면: sh $H/go f)"
  fi
  exit 1
fi

# 1) AGENTS.md — 과제가 준 AGENTS.md가 있으면 지우지 않고 맨 아래에 붙인다(과제 규칙이 우선)
if [ -f AGENTS.md ] && grep -q "$MARK" AGENTS.md; then
  say "AGENTS.md: 하네스 규칙이 이미 있어 그대로 둠"
elif [ -f AGENTS.md ]; then
  printf '\n' >> AGENTS.md
  cat "$H/AGENTS.md" >> AGENTS.md
  say "AGENTS.md: 과제 규칙 아래에 하네스 규칙을 붙임 (과제 규칙이 우선)"
else
  cat "$H/AGENTS.md" > AGENTS.md
  say "AGENTS.md: 새로 만듦"
fi

# 2) MSG.md — 내 메모. 이미 있으면 덮어쓰지 않는다
if [ -f MSG.md ] && grep -q '^# MSG.md' MSG.md; then
  say "MSG.md: 이미 있어 그대로 둠 (새로 받으려면: cat $H/MSG.md > MSG.md)"
elif [ -f MSG.md ]; then
  say "MSG.md: 과제에 같은 이름 파일이 있어 건드리지 않음 → 내 메모는 $H/MSG.md"
else
  cat "$H/MSG.md" > MSG.md
  say "MSG.md: 만듦 (내 메모 · 커밋 안 함)"
fi

# 3) 다음 할 일
say "다음: 같은 탭에 codex → 'Update available'이면 2 · 'Do you trust'면 Enter → /model Sol 6.1 medium → /status → 정찰"
say "흐름 · 문장은 편집기에서 MSG.md (맨 위 '흐름')"
