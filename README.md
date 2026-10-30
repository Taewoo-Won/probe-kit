# probe-kit

## 작업 하네스 (AGENTS.md · MSG.md)

Codex CLI로 업무 담당자와 협업하며 과제를 풀 때 쓰는 개인 작업 규칙과 문장 템플릿입니다. 과제의 답이나 코드는 들어 있지 않습니다.

- `AGENTS.md` — 한 단어 명령(정찰 · 초안 · 반영 · 계획 · 다음 · 마무리)의 정의와 보고 형식
- `MSG.md` — 업무 담당자에게 보내는 메시지와 Codex에 치는 문장 모음

과제 폴더의 터미널에서 (작업 폴더 밖인 /tmp/harness에 받는다):

```
git clone https://github.com/Taewoo-Won/probe-kit /tmp/harness
cat /tmp/harness/AGENTS.md >> AGENTS.md
cat /tmp/harness/MSG.md > MSG.md
```

- 둘째 줄은 `>>`(맨 아래에 덧붙이기)라서 과제 폴더에 AGENTS.md가 이미 있어도 지워지지 않는다. 원래 있었으면 원래 규칙과 부딪히는 줄을 지운다(과제 작성자 규칙이 우선). 한 번만 친다.
- MSG.md는 편집기에서 열어 보는 메모다. 맨 위에 흐름(남은 시간별 할 일)과 끼어들 때 할 일, 그 아래에 복사해 쓰는 문장이 있다. AGENTS.md가 Codex에게 읽지도 커밋하지도 말라고 정해 둔다.

연습용 키트는 `kit` 브랜치에 있다.
