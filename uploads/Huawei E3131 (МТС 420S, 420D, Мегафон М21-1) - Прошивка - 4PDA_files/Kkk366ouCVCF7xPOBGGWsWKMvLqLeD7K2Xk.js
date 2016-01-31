var ipsmenu=window.ipsmenu=new ips_menu;function ips_menu(){var t=this;t.menu_registered=[];t.menu_openfuncs=[];t.menu_over_css=[];t.menu_out_css=[];t.menu_open_event=[];t.dynamic_register=[];t.menu_cur_open=null;t.dynamic_html=null;t.guid=0;return t}
ips_menu.prototype.register=function(cid,callback,menu_over_css,menu_out_css,event_type){var t=this;if(event_type)t.menu_open_event[cid]=event_type=="onmouseover"?"onmouseover":"onclick";else event_type="onclick";t.menu_registered[cid]=new ips_menu_class(cid);if(callback)t.menu_openfuncs[cid]=callback;if(menu_over_css&&menu_out_css){t.menu_over_css[cid]=menu_over_css;t.menu_out_css[cid]=menu_out_css}return t.menu_registered[cid]};ips_menu.prototype.close=function(){var t=this;if(t.menu_cur_open)t.menu_registered[t.menu_cur_open].close()};
function ips_menu_class(cid){var t=this;t.cid=cid;t.initialized=false;(function(){function r(f){/in/.test(document.readyState)?setTimeout(function(){r(f)},9):f()}r(function(){t.init_control_object();t.init_menu()})})();return t}ips_menu.prototype.commonHandle=function(ev){var hs=this.events[ev.type];for(var g in hs){var h=hs[g],r=h.call(this,ev);if(r===false){ev.preventDefault&&ev.preventDefault();ev.returnValue=false;ev.stopPropagation&&ev.stopPropagaton();ev.cancelBubble=true}}};
ips_menu_class.prototype.addEvent=function(o,e,f){if(o.addEventListener)o.addEventListener(e,f,false);else if(o.attachEvent){if(!f.guid)f.guid=++ipsmenu.guid;if(!o.events){o.events={};o.handler=function(ev){if(typeof ev!=="undefined")return ipsmenu.commonHandle.call(o,ev)}}if(!o.events[e]){o.events[e]={};o.attachEvent("on"+e,o.handler)}o.events[e][f.guid]=f}else o["on"+e]=f};
ips_menu_class.prototype.init_control_object=function(){var t=this;t.cid_obj=document.getElementById(t.cid);if(!t.cid_obj)return;try{t.cid_obj.style.cursor="pointer"}catch(e){t.cid_obj.style.cursor="hand"}t.cid_obj.unselectable=true;if(ipsmenu.menu_open_event[t.cid]=="onmouseover")t.addEvent(t.cid_obj,"mouseover",ips_menu_events.prototype.event_onclick);else{t.addEvent(t.cid_obj,"click",ips_menu_events.prototype.event_onclick);t.addEvent(t.cid_obj,"mouseover",ips_menu_events.prototype.event_onmouseover)}t.addEvent(t.cid_obj,
"mouseout",ips_menu_events.prototype.event_onmouseout)};ips_menu_class.prototype.init_menu=function(){var t=this;t.cid_menu_obj=document.getElementById(t.cid+"_menu");if(t.cid_menu_obj){if(t.initialized)return;t.cid_menu_obj.style.display="none";t.cid_menu_obj.style.position="absolute";t.cid_menu_obj.style.left="0px";t.cid_menu_obj.style.top="0px";t.addEvent(t.cid_menu_obj,"click",ipsclass.cancel_bubble_low);t.cid_menu_obj.zIndex=50;t.initialized=true}};
ips_menu_class.prototype.open=function(obj){var t=this;if(!t.cid_menu_obj){t.initialized=false;t.init_menu()}if(ipsmenu.menu_cur_open!=null)ipsmenu.menu_registered[ipsmenu.menu_cur_open].close();if(ipsmenu.menu_cur_open==obj.id)return false;ipsmenu.menu_cur_open=obj.id;var left_px=ipsclass.get_obj_leftpos(obj);var top_px=ipsclass.get_obj_toppos(obj)+obj.offsetHeight;var ifid=obj.id;t.cid_menu_obj.style.zIndex=-1;t.cid_menu_obj.style.display="";var width=parseInt(t.cid_menu_obj.style.width)?parseInt(t.cid_menu_obj.style.width):
t.cid_menu_obj.offsetWidth;if(left_px+width>=document.body.clientWidth)left_px=left_px+obj.offsetWidth-width;if(is_moz)top_px-=1;t.cid_menu_obj.style.left=left_px+"px";t.cid_menu_obj.style.top=top_px+"px";t.cid_menu_obj.style.zIndex=100;if(ipsmenu.menu_openfuncs[obj.id])eval(ipsmenu.menu_openfuncs[obj.id]);if(is_ie)try{if(!document.getElementById("if_"+obj.id)){var iframeobj=document.createElement("iframe");iframeobj.src="javascript:;";iframeobj.id="if_"+obj.id;document.getElementsByTagName("body").appendChild(iframeobj)}else var iframeobj=
document.getElementById("if_"+obj.id);iframeobj.scrolling="no";iframeobj.frameborder="no";iframeobj.className="iframeshim";iframeobj.style.position="absolute";iframeobj.style.width=parseInt(t.cid_menu_obj.offsetWidth)+"px";iframeobj.style.height=parseInt(t.cid_menu_obj.offsetHeight)+"px";iframeobj.style.top=t.cid_menu_obj.style.top;iframeobj.style.left=t.cid_menu_obj.style.left;iframeobj.style.zIndex=99;iframeobj.style.display="block"}catch(error){}if(t.cid_obj.editor_id){t.cid_obj.state=true;IPS_editor[t.cid_obj.editor_id].set_menu_context(t.cid_obj,
"mousedown")}return false};
ips_menu_class.prototype.close=function(){var t=this;if(t.cid_menu_obj!=null)t.cid_menu_obj.style.display="none";else if(ipsmenu.menu_cur_open!=null)ipsmenu.menu_registered[ipsmenu.menu_cur_open].cid_menu_obj.style.display="none";ipsmenu.menu_cur_open=null;if(t.cid_obj)if(ipsmenu.menu_out_css[t.cid_obj.id])t.cid_obj.className=ipsmenu.menu_out_css[t.cid_obj.id];if(is_ie)try{document.getElementById("if_"+t.cid).style.display="none"}catch(error){}if(t.cid_obj.editor_id){t.cid_obj.state=false;IPS_editor[t.cid_obj.editor_id].set_menu_context(t.cid_obj,
"mouseout")}};ips_menu_class.prototype.hover=function(e){var t=this;if(ipsmenu.menu_cur_open!=null)if(ipsmenu.menu_registered[ipsmenu.menu_cur_open].cid!=t.id)t.open(e)};function ips_menu_events(){}ips_menu_events.prototype.event_safari_onclick_handler=function(){if(this.id)window.location=document.getElementById(this.id).href};
ips_menu_events.prototype.event_onmouseover=function(e){var t=this;e=ipsclass.cancel_bubble(e,true);ipsmenu.menu_registered[t.id].hover(t);if(ipsmenu.menu_over_css[t.id])t.className=ipsmenu.menu_over_css[t.id]};ips_menu_events.prototype.event_onmouseout=function(e){var t=this;e=ipsclass.cancel_bubble(e,true);if(ipsmenu.menu_out_css[t.id]&&ipsmenu.menu_cur_open!=t.id)t.className=ipsmenu.menu_out_css[t.id]};
ips_menu_events.prototype.event_onclick=function(e){var t=this;e=ipsclass.cancel_bubble(e,true);if(ipsmenu.menu_cur_open==null){if(ipsmenu.menu_over_css[t.id])t.className=ipsmenu.menu_over_css[t.id];ipsmenu.menu_registered[t.id].open(t)}else if(ipsmenu.menu_cur_open==t.id){ipsmenu.menu_registered[t.id].close();if(ipsmenu.menu_out_css[t.id])t.className=ipsmenu.menu_out_css[t.id]}else{if(ipsmenu.menu_over_css[t.id])t.className=ipsmenu.menu_over_css[t.id];ipsmenu.menu_registered[t.id].open(t)}};
function menu_do_global_init(){document.onclick=menu_action_close;if(ipsmenu.dynamic_html!=null&&ipsmenu.dynamic_html!="");if(ipsmenu.dynamic_register.length)for(var i=0;i<ipsmenu.dynamic_register.length;i++)if(ipsmenu.dynamic_register[i])ipsmenu.register(ipsmenu.dynamic_register[i])}function menu_action_close(e){try{if(e.button==2||e.button==3)return}catch(acold){}ipsmenu.close(e)};
function nav_menu_titles(){
document.write('<a id="usermenu-0" class="usermenu" unselectable="true" href="/forum/?showforum=281" title="Android">Android</a>\
&nbsp; \
<a id="usermenu-1" class="usermenu" unselectable="true" href="/forum/?showforum=201" title="Windows Mobile">WM</a>\
&nbsp; \
<a id="usermenu-2" class="usermenu" unselectable="true" href="/forum/?showforum=355" title="Windows Phone">WP</a>\
&nbsp; \
<a id="usermenu-3" class="usermenu" unselectable="true" href="/forum/?showforum=295" title="iOS">iOS</a>\
&nbsp; \
<a id="usermenu-4" class="usermenu" unselectable="true" href="/forum/?showforum=444" title="Symbian">Symbian</a>\
&nbsp; \
<a id="usermenu-5" class="usermenu" unselectable="true" href="/forum/?showforum=442" title="������ ���������">������</a>\
&nbsp; \
<a id="usermenu-6" class="usermenu" unselectable="true" href="/forum/#" title="�������� ������">��������</a>\
&nbsp; \
<a id="usermenu-7" class="usermenu" unselectable="true" href="http://4pda.ru/devdb/" title="DevDB - ������� ���������" target="_blank">����������</a>\
&nbsp; \
<a class="usermenu" unselectable="true" href="http://devfaq.ru/" title="DevFAQ - ���� ������ �� �����������" target="_blank">���� ������</a>\
');
}
function nav_menu_body(){
document.write('<div class="upopupmenu-new" id="usermenu-0_menu" style="display: none; width: 130px; position: absolute; left: 0px; top:0px;">\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=282" title="������ ������ �� Android">������ ������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=283" title="FAQ �� Android">FAQ</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showtopic=112220" title="��������� ��� Android">���������</a></div>\
<div class="upopupmenu-item-last"><a href="/forum/index.php?showtopic=117270" title="���� ��� Android">����</a></div>\
</div>\
<div class="upopupmenu-new" id="usermenu-1_menu" style="display: none; width: 130px; position: absolute; left: 0px; top:0px;">\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=8" title="������ ������ �� Windows Mobile">������ ������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=252" title="FAQ �� Windows Mobile">FAQ</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showtopic=20758" title="��������� ��� Windows Mobile">���������</a></div>\
<div class="upopupmenu-item-last"><a href="/forum/index.php?showtopic=19674" title="���� ��� Windows Mobile">����</a></div>\
</div>\
<div class="upopupmenu-new" id="usermenu-2_menu" style="display: none; width: 130px; position: absolute; left: 0px; top:0px;">\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=361" title="������ ������ �� Windows Phone">������ ������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=393" title="FAQ �� Windows Phone">FAQ</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showtopic=236428" title="��������� ��� Windows Phone">���������</a></div>\
<div class="upopupmenu-item-last"><a href="/forum/index.php?showtopic=321019" title="���� ��� Windows Phone">����</a></div>\
</div>\
<div class="upopupmenu-new" id="usermenu-3_menu" style="display: none; width: 140px; position: absolute; left: 0px; top:0px;">\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=139" title="���������� Apple">����������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=142" title="������ ������ �� iPod Touch, iPhone � iPad">������ ������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=144" title="FAQ �� iPod Touch, iPhone iPad">FAQ</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=140" title="��������� ��� iPod Touch, iPhone b iPad">���������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=141" title="���� ��� iPod Touch, iPhone � iPad">����</a></div>\
</div>\
<div class="upopupmenu-new" id="usermenu-4_menu" style="display: none; width: 130px; position: absolute; left: 0px; top:0px;">\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=54" title="������ ������ �� Symbian">������ ������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=463" title="FAQ �� Symbian">FAQ</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showtopic=106137" title="��������� ��� Symbian">���������</a></div>\
<div class="upopupmenu-item-last"><a href="/forum/index.php?showtopic=107144" title="���� ��� Symbian">����</a></div>\
</div>\
<div class="upopupmenu-new" id="usermenu-5_menu" style="display: none; width: 130px; position: absolute; left: 0px; top:0px;">\
<div class="upopupmenu-category"><a href="/forum/index.php?showforum=443" title="WM Smartphones">WM Smartphones</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=464" title="FAQ �� WM Smartphones">FAQ</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showtopic=158429" title="��������� ��� WM Smartphones">���������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showtopic=157783" title="���� ��� WM Smartphones">����</a></div>\
<div class="upopupmenu-category"><a href="/forum/index.php?showforum=447" title="Maemo/MeeGo">Maemo/MeeGo</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=454" title="��������� ��� Maemo/MeeGo">���������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=455" title="���� ��� Maemo/MeeGo">����</a></div>\
<div class="upopupmenu-category"><a href="/forum/index.php?showforum=446" title="Bada">Bada</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=461" title="FAQ �� Bada">FAQ</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=365" title="��������� ��� Bada">���������</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=366" title="���� ��� Bada">����</a></div>\
<div class="upopupmenu-category"><a href="/forum/index.php?showforum=446" title="��������� � �������� �� Java">Java</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=308" title="FAQ �� ���������� � ��������� �� Java">FAQ</a></div>\
<div class="upopupmenu-item"><a href="/forum/index.php?showforum=196" title="��������� ��� ���������� � ��������� �� Java">���������</a></div>\
<div class="upopupmenu-item-last"><a href="/forum/index.php?showforum=196" title="���� ��� ���������� � ��������� �� Java">����</a></div>\
</div>\
<div class="upopupmenu-new" id="usermenu-6_menu" style="display: none; width: 140px; position: absolute; left: 0px; top:0px;">\
<div class="upopupmenu-item"><a href="http://4pda.ru/forum/index.php?showtopic=121191" title="10 ��������� �������">10 ��������� �������</a></div>\
<div class="upopupmenu-item"><a href="http://4pda.ru/forum/index.php?showtopic=89878" title="FAQ �� ������">FAQ �� ������</a></div>\
<div class="upopupmenu-item-last"><a href="http://4pda.ru/forum/index.php?showforum=373" title="�������� �� ������">�������� �� ������</a></div>\
</div>\
<div class="upopupmenu-new" id="usermenu-7_menu" style="display: none; width: 130px; position: absolute; left: 0px; top:0px;">\
<div class="upopupmenu-item"><a href="http://4pda.ru/devdb/phones/" target="_blank" title="�������������">��������</a></div>\
<div class="upopupmenu-item"><a href="http://4pda.ru/devdb/pad/" target="_blank" title="��������">��������</a></div>\
<div class="upopupmenu-item"><a href="http://4pda.ru/devdb/ebook/" target="_blank" title="����������� �����">����������� �����</a></div>\
<div class="upopupmenu-item-last"><a href="http://4pda.ru/devdb/smartwatch/" target="_blank" title="����� ����">����� ����</a></div>\
</div>\
');
for(var i=0;i<9;i++)
{
	var o=document.getElementById('usermenu-'+i);
	if(o)
	{
//		o.onclick=function(){return false;};
		ipsmenu.register("usermenu-"+i,"","usermenu","usermenu_out");
	}
}
}
