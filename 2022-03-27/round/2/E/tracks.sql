SELECT InvoiceLine.TrackId       AS TrackId,
       SUM(InvoiceLine.Quantity) AS Total
FROM InvoiceLine
           INNER JOIN
     Invoice
     ON
                 InvoiceLine.InvoiceId = Invoice.InvoiceId
WHERE Invoice.InvoiceDate LIKE '201%' and InvoiceLine.TrackId in (select TrackId from Track )
GROUP BY InvoiceLine.TrackId
ORDER BY Total DESC, InvoiceLine.TrackId;