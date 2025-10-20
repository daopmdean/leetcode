-- Write your PostgreSQL query statement below

SELECT name as employee
FROM Employee e
WHERE e.managerId IS NOT null 
AND e.salary > (SELECT m.salary
            FROM Employee m
            WHERE m.id = e.managerId)

