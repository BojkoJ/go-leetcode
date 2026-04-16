# go-leetcode

Repozitář s řešeními úloh z LeetCode v jazyce Go. Každá úloha má vlastní složku (pojmenovanou podle ID a názvu) a uvnitř jednoduchý spustitelný příklad.

## Struktura

- Každá úloha je ve vlastní složce, např. `1-two-sum/`
- Uvnitř typicky:
  - `main.go` – implementace + krátký runner
  - `go.mod` – samostatný modul pro danou úlohu (jednoduché spouštění izolovaně)

## Spuštění

Předpoklady: Go nainstalované lokálně.

Spuštění konkrétní úlohy:

```bash
cd 1-two-sum
go run .
```

## Log vyřešených úloh

Záznamy jsou řazené podle čísla úlohy.

| # | Název | Složka | Technika | Složitost | Status | Poznámka |
|---:|------|--------|----------|-----------|--------|----------|
| 1 | Two Sum | `1-two-sum/` | Hash map (doplněk), brute force | O(n) / O(n²) | done | 2 řešení: (1) dvojitý loop; (2) průchod polem + mapa `hodnota → index` pro rychlé ověření, zda existuje doplněk `target - num`. |
| 2 | Add Two Numbers | `2-add-two-numbers/` | Linked list, dummy head, carry | O(max(m,n)) / O(max(m,n)) | done | Sčítání po cifrách v reverse pořadí, průchod oběma listy + přenos `carry`; stavba výsledku přes „dummy head“ uzel. |
| 3 | Longest Substring Without Repeating Characters | `3-longest-substring-without-repeating-characters/` | Sliding window, Hash map | O(n) / O(min(m, a)) | done | Použití hash mapy k uchovávání posledních indexů znaků a "sliding window" k rychlému hledání bez zbytečného procházení všech možností. |
