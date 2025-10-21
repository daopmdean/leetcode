-- Write your PostgreSQL query statement below

SELECT c.name as customers
FROM Customers c
LEFT JOIN Orders o
ON c.id = o.customerId
WHERE o.id IS null
