package clickhouse

import "fmt"

const (
	qSamples = `
SELECT 
    StackHash AS hash, 
    sum(Value) AS value
FROM @@table_profiling_samples@@
WHERE 
    ServiceName IN (@service) AND 
    Type = @type AND 
    Start < @to AND End > @from AND
    (empty(@containers) OR has(@containers, Labels['container.id']))
GROUP BY StackHash
`
	qSamplesDiff = `
SELECT 
	StackHash AS hash,
	sum(CASE WHEN End < @middle THEN Value ELSE 0 END) AS base,
	sum(CASE WHEN Start > @middle THEN Value ELSE 0 END) AS comp
FROM @@table_profiling_samples@@
WHERE 
    ServiceName IN (@service) AND 
    Type = @type AND 
    Start < @to AND End > @from AND
    (empty(@containers) OR has(@containers, Labels['container.id']))
GROUP BY StackHash
`
	qStacks = `
SELECT 
    Hash AS hash, 
    any(Stack) AS stack
FROM @@table_profiling_stacks@@
WHERE 
    ServiceName IN (@service) AND 
    LastSeen > @from
GROUP BY Hash
`
	qProfiles = `
SELECT 
    count(distinct Start) AS count
FROM @@table_profiling_samples@@
WHERE 
    ServiceName IN (@service) AND 
    Type = @type AND 
    Start < @to AND End > @from AND
    (empty(@containers) OR has(@containers, Labels['container.id']))
`
)

var (
	qProfileTypes = "SELECT DISTINCT ServiceName, Type FROM @@table_profiling_profiles@@ WHERE LastSeen >= @from"
	qProfile      = fmt.Sprintf("WITH samples AS (%s), stacks AS (%s) SELECT value, stack FROM samples JOIN stacks USING(hash)", qSamples, qStacks)
	qProfileAvg   = fmt.Sprintf("WITH samples AS (%s), stacks AS (%s), profiles AS (%s) SELECT toInt64(value/count), stack FROM samples JOIN stacks USING(hash), profiles", qSamples, qStacks, qProfiles)
	qProfileDiff  = fmt.Sprintf("WITH samples AS (%s), stacks AS (%s) SELECT base, comp, stack FROM samples JOIN stacks USING(hash)", qSamplesDiff, qStacks)
)
