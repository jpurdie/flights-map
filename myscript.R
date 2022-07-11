library(maps)
library(geosphere)

map("world", col="#191919", fill=TRUE, bg="#000000")

airports <- read.csv("my-aiports-data.csv", header=TRUE) 
flights <- read.csv("my-flights-data.csv", header=TRUE, as.is=TRUE)

pal <- colorRampPalette(c("#d2d2d2", "black"))
colors <- pal(100)
 
fsub <- flights[flights$,]
fsub <- fsub[order(fsub$cnt),]
maxcnt <- max(fsub$cnt)

print(maxcnt)

for (j in 1:length(fsub$airline)) {
    air1 <- airports[airports$iata == fsub[j,]$airport1,]
    air2 <- airports[airports$iata == fsub[j,]$airport2,]
     
    inter <- gcIntermediate(c(air1[1,]$long, air1[1,]$lat), c(air2[1,]$long, air2[1,]$lat), n=100, addStartEnd=TRUE)
    colindex <- round( (fsub[j,]$cnt / maxcnt) * length(colors) )
             
    lines(inter, col=colors[colindex], lwd=0.8)
}



