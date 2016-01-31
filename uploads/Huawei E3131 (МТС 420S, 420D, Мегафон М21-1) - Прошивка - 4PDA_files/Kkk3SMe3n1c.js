var e=["http:\/\/4pda.ru\/","http:\/\/4pda.to\/","http:\/\/4pda.ws\/","http:\/\/s.4pda.to\/","http:\/\/oxn.gerkon.eu\/","https:\/\/adclick.g.doubleclick.net\/","http:\/\/ad.doubleclick.net\/"]
var i=["http:\/\/4pda.ru\/about\/our-projects\/","http:\/\/4pda.ru\/advert\/for-developers\/"]
var x,y,z;
var c=(document.documentElement?document.documentElement:document.body).getElementsByTagName("a");
for(x=0;x<c.length;++x)
{
	var h=c[x].href,t;
	if(!h||(h.indexOf("http")!=0)||(c[x].getAttribute('data-notrack'))){continue;}
	for(y=0;y<e.length;++y){if(h.indexOf(e[y])==0){break;}}
	if(y<e.length)
	{
		for(z=0;z<i.length;++z){if(h.indexOf(i[z])==0){break;}}
		if(z==i.length){continue;}
	}
	h="http://4pda.ru/pages/go/?u="+encodeURIComponent(h);
	if(t=c[x].getAttribute('data-trackmark')){h+='&m='+t;}
	c[x].href=h;
}
