select Employee.FirstName || ' ' || Employee.LastName as Name, count(Invoiceid) as count
from Invoice
         INNER join Customer on Customer.CustomerId = Invoice.CustomerId
         INNER join Employee on Employee.EmployeeId = Customer.SupportRepId
WHERE Invoice.InvoiceDate LIKE "201%"
group by Customer.SupportRepId
order by count desc
limit 3;
