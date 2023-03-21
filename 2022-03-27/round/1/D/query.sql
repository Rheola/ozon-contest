select BillingCountry,  SUM(Total)  as total
from Invoice
group by BillingCountry
order by total;
