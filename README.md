# probe-kit

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
