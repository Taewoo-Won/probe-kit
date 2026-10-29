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
- MSG.md는 편집기에서 열어 문장을 복사하는 메모다. AGENTS.md가 Codex에게 읽지도 커밋하지도 말라고 정해 둔다.

## 모의고사 연습 키트 (probe-kit.js)

AI 협업 코딩 모의고사를 **실제 Codex CLI**로 연습할 때 쓰는 키트. 가상 과제(README)와 가상 데이터, 일부 과제의 가상 '전임자 코드', 제출용 pack 스크립트만 들어 있어요. 실제 회사·시험 문제와는 관계없어요.

## 받기

```
curl -fsSL https://raw.githubusercontent.com/Taewoo-Won/probe-kit/main/probe-kit.js -o ~/probe-kit.js
```

## 쓰기

```
node ~/probe-kit.js list              # 과제 목록
node ~/probe-kit.js new recipe        # 과제 폴더 만들기 (빈 상태: README·데이터만)
cd "$(node ~/probe-kit.js where)" && codex
node ~/probe-kit.js pack              # 끝나면 그 폴더에서 → PROBE1 줄을 모의고사 제출 칸에 붙여넣기
```

- `--b` : 합의문(AGREEMENT.md)만 깔린 상태로 시작 (연습 B)
- 봉인 과제는 `--open` 을 붙여야 열려요
- Node.js 18 이상, 외부 패키지 없음
- `pack`은 과제 폴더의 파일(제공 데이터는 바뀌었을 때만), git 기록, 이 폴더에서 돌린 Codex 세션의 지시·명령·토큰, README에 적힌 `node` 실행 결과를 한 줄로 묶어요
