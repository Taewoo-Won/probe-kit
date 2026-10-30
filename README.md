# probe-kit

시작 한 줄 — 과제 폴더(README.md가 있는 곳)의 터미널에서:

```
git clone https://github.com/Taewoo-Won/probe-kit /tmp/harness; sh /tmp/harness/go
```

## 작업 하네스 (AGENTS.md · MSG.md)

Codex CLI로 업무 담당자와 협업하며 과제를 풀 때 쓰는 개인 작업 규칙과 문장 템플릿입니다. 과제의 답이나 코드는 들어 있지 않습니다.

- `AGENTS.md` — 한 단어 명령(정찰 · 초안 · 반영 · 계획 · 다음 · 마무리)의 정의와 보고 형식
- `MSG.md` — 업무 담당자에게 보내는 메시지와 Codex에 치는 문장 모음
- `go` — 시작 스크립트. 하는 일은 셋뿐이다(작업 폴더 안에 잘못 받은 하네스 사본이 있으면 지운다).
  1. AGENTS.md 맨 아래에 하네스 규칙을 붙인다. 과제가 준 AGENTS.md가 있으면 지우지 않고 그 아래에 붙인다(과제 규칙이 우선). 이미 붙어 있으면 건너뛴다.
  2. MSG.md를 만든다. 이미 있으면 덮어쓰지 않는다.
  3. 다음 할 일(codex → 팝업 → /model → /status)을 보여 준다.

  설치 · 네트워크 · git 커밋 · Codex 설정은 건드리지 않는다. 여러 번 쳐도 결과가 같아서 한 줄을 통째로 다시 쳐도 된다. 하네스는 작업 폴더 밖인 /tmp/harness에 받는다.

## 막힐 때

- `already exists`가 나와도 괜찮다: 이미 받은 것이고, `;` 뒤의 go는 그대로 돈다.
- `README가 없어`가 나온다: go가 알려 주는 대로 과제 폴더로 옮긴 뒤 `sh /tmp/harness/go`. 여기가 과제 폴더가 맞으면 `sh /tmp/harness/go f`.
- `cannot open /tmp/harness/go`: clone이 안 된 것이다. 주소(https:// · probe-kit 다음 한 칸 · /tmp)를 확인하고 한 줄을 다시 친다.
- go가 안 되면 손으로 같은 일을 한다:
  ```
  cat /tmp/harness/AGENTS.md >> AGENTS.md
  cat /tmp/harness/MSG.md > MSG.md
  ```
  첫 줄은 `>>`(맨 아래에 덧붙이기)라서 과제 AGENTS.md가 지워지지 않는다. 한 번만 친다.
- MSG.md가 망가졌다: `cat /tmp/harness/MSG.md > MSG.md`
- 하네스를 최신으로: `git -C /tmp/harness pull`

MSG.md는 편집기에서 열어 보는 메모다. 맨 위에 흐름(남은 시간별 할 일)과 끼어들 때 할 일, 그 아래에 복사해 쓰는 문장이 있다. AGENTS.md가 Codex에게 읽지도 커밋하지도 말라고 정해 둔다.

연습용 키트는 `kit` 브랜치에 있다.
