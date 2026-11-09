#!/bin/bash
# probe-kit 결선 모드 실제 Codex 시험: 정찰 → (다음 할 일 따라) 4번까지
T=/root/pk_selftest
CODEX=$(command -v codex || echo /root/.local/bin/codex)
rm -rf "$T"; mkdir -p "$T/task/data"; cd "$T" || exit 1
git clone -q https://github.com/Taewoo-Won/probe-kit h || { echo "clone 실패" > "$T/summary.txt"; exit 1; }
cd task
cat > README.md <<'EOF'
# 식각 공정 수율 하락 — 무엇을 만들지는 여러분이 정합니다

## 상황
메모리 팹 식각(Etch) 공정에서 최근 4주 동안 제품 A의 웨이퍼 가장자리(edge) 불량이 늘어 수율이 떨어졌습니다.
공정 엔지니어는 매주 월요일 엑셀로 로트별 공정 이력을 한 줄씩 대조해 원인 장비를 찾습니다. 한 번에 약 6시간이 걸리고, 원인을 찾는 동안에도 불량 로트는 계속 나옵니다.

## 데이터 (data/)
- lot_history.csv — lot_id, step, tool, chamber, time_in
- lot_yield.csv — lot_id, week, yield_pct, edge_defect_pct
- pm_log.csv — tool, chamber, pm_time, pm_type

## 제출
- 이 폴더에 만든 결과물
- SUBMISSION.md : 무엇을 왜 만들었고, 어떻게 확인하는지

정답은 정해져 있지 않습니다. 현장이 당연하게 여겨온 방식을 새로운 시각으로 보고, AI로 해결의 가능성을 보여 주세요.
EOF
python3 - <<'EOF'
import random, csv, datetime as dt
random.seed(7)
tools={'ETCH':['ETCH-01','ETCH-02','ETCH-03'],'DEP':['DEP-01','DEP-02'],'CMP':['CMP-01','CMP-02']}
start=dt.datetime(2026,9,7,6,0); hist=[]; yld=[]
for i in range(400):
    lot=f"LOT{i:04d}"; t=start+dt.timedelta(hours=i*1.68); week=1+int((t-start).days//7); path={}
    for k,step in enumerate(['DEP','ETCH','CMP']):
        tool=random.choice(tools[step]); ch=random.choice('ABC'); path[step]=(tool,ch)
        hist.append([lot if random.random()>0.03 else lot.lower(),step,tool,ch,(t+dt.timedelta(hours=k*3)).strftime('%Y-%m-%d %H:%M')])
    y=random.gauss(92,1.5); e=max(0,random.gauss(1.2,0.4))
    if path['ETCH']==('ETCH-03','C') and week>=3: e+=random.gauss(4.0,0.8); y-=random.gauss(3.5,0.7)
    yld.append([lot,week,round(y,2),round(e,2)])
hist+=[hist[10],hist[200]]
for name,head,rows in [('lot_history',['lot_id','step','tool','chamber','time_in'],hist),('lot_yield',['lot_id','week','yield_pct','edge_defect_pct'],yld),
  ('pm_log',['tool','chamber','pm_time','pm_type'],[['ETCH-01','A','2026-09-10 08:00','정기'],['ETCH-03','C','2026-09-21 14:00','정기'],['ETCH-02','B','2026-09-24 09:00','정기'],['DEP-01','A','2026-09-15 10:00','점검'],['ETCH-03','C','2026-09-28 13:00','긴급']])]:
    with open(f'data/{name}.csv','w',newline='') as f: w=csv.writer(f); w.writerow(head); w.writerows(rows)
EOF
cp ../h/AGENTS.md AGENTS.md
git init -q && git add -A && git -c user.name=tester -c user.email=t@t commit -qm "과제 원본"
LOG="$T/log.txt"; : > "$LOG"; next="정찰"; steps=""
for i in 1 2 3 4; do
  echo "=== [$i] $next ===" >> "$LOG"
  timeout 1500 "$CODEX" exec --skip-git-repo-check --sandbox workspace-write "$next" < /dev/null >> "$LOG" 2>&1
  nx=$(grep -o '다음 할 일: *[가-힣]*' "$LOG" | tail -1 | sed 's/다음 할 일: *//')
  steps="$steps $next→${nx:-?}"
  [ -z "$nx" ] && break
  next="$nx"
done
{
echo "흐름:$steps"
grep -m2 -E '과제 모양|담당자:' WORKLOG.md 2>/dev/null | cut -c1-110
echo "합의문 세 칸: $(grep -c -E '지금 방식|질문 셋|고른 질문' AGREEMENT.md 2>/dev/null)"
grep -m1 -A1 -E '고른 질문' AGREEMENT.md 2>/dev/null | cut -c1-140
grep -E '^\s*[0-9]+[.)]' PLAN.md 2>/dev/null | head -7 | cut -c1-80
echo "커밋: $(git log --oneline | wc -l)개"
echo "Usage 오류: $(grep -c 'Usage:' "$LOG")"
} > "$T/summary.txt"
echo 끝 >> "$T/summary.txt"
