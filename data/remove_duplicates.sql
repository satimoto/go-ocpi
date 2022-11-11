DELETE FROM locations where id IN (SELECT id
  FROM (SELECT id, 
	ROW_NUMBER() OVER (partition BY uid ORDER BY id) AS rn
    FROM locations) AS l
WHERE l.rn > 1)

DELETE FROM evses where id IN (SELECT id
  FROM (SELECT id, 
	ROW_NUMBER() OVER (partition BY uid ORDER BY id) AS rn
    FROM evses) AS e
WHERE e.rn > 1)