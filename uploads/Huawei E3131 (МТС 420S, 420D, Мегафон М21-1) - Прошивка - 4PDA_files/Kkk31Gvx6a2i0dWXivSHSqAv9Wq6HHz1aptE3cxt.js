var ipsattach_css={
'css_menu_row' : 'attach-menu-row'
,'css_menu_row_over' : 'attach-menu-row-over'
};
//var ipsattach=new ips_attach();

function ips_attach(edid) {
	this.add_app_id = '';
	edid&&(this.edid = edid);
	this.iframe_id = 'ips--iframe-obj-attach';
	this.iframe_obj = null;
	this.iframe_init_url = '';
	this.iframe_remove_url = '';
	this.iframe_add_to_div = 'ips-attach-div-iframe';
	this.iframe_add_to_div_obj = null;
	this.iframe_parent_div = 'ips-attach-div-parent';
	this.iframe_parent_div_obj = null;
	this.status_div = 'ips-attach-div-status';
	this.status_div_obj = null;
	this.dropdown_text = 'ips-attach-menu-text';
	this.dropdown_text_obj = null;
	this.iframe_height = 300;
	this.iframe_classname = 'attach-iframe';
	this.message_div_id = 'ips-attach-message';
	this.message_div_obj = null;
	this.message_span_id = 'ips-attach-message-span';
	this.message_span_obj = null;
	this.status_msg = '';
	this.status_is_error = '';
	this.current_items = {};//new Array();
	this.current_item_menu = 'ips-attach-menu';
	this.current_item_menu_obj = null;
	this.images_url = null;
	this.lang = new Array();
	this.css_menu_row = ipsattach_css.css_menu_row;
	this.css_menu_row_over = ipsattach_css.css_menu_row_over;
	this.show_attach_menu = 1;
	this.can_add_filecode = 1;
	this.onload_function = '';
	this.has_upload_pending = 0;
	this.frm_filelist = null;
	this.load_unlinked = '';
	this.init = function () {
		if (this.add_app_id) {
			this.iframe_id = 'ips--iframe-obj-attach' + this.add_app_id;
			this.iframe_add_to_div = 'ips-attach-div-iframe' + this.add_app_id;
			this.iframe_parent_div = 'ips-attach-div-parent' + this.add_app_id;
			this.status_div = 'ips-attach-div-status' + this.add_app_id;
			this.dropdown_text = 'ips-attach-menu-text' + this.add_app_id;
			this.message_div_id = 'ips-attach-message' + this.add_app_id;
			this.message_span_id = 'ips-attach-message-span' + this.add_app_id;
			this.current_item_menu = 'ips-attach-menu' + this.add_app_id;
		}
	var tmpid='file-list'+(this.add_app_id?'['+this.add_app_id+']':'');
	document.write('<input type="hidden" name="'+tmpid+'" id="'+tmpid+'" />');
	if (''!=this.load_unlinked)this.iframe_init_url+='&load_unlinked='+this.load_unlinked;
	this.frm_filelist=document.getElementById(tmpid);
		this.iframe_add_to_div_obj = document.getElementById(this.iframe_add_to_div);
		this.iframe_parent_div_obj = document.getElementById(this.iframe_parent_div);
		this.message_div_obj = document.getElementById(this.message_div_id);
		this.message_span_obj = document.getElementById(this.message_span_id);
		this.status_div_obj = document.getElementById(this.status_div);
		this.dropdown_text_obj = document.getElementById(this.dropdown_text);
		if (this.show_attach_menu) {
			this.current_item_menu_obj = document.getElementById(this.current_item_menu);
			this.current_item_menu_obj.add_app_id = (this.add_app_id) ? this.add_app_id : '';
			ipsmenu.register(this.current_item_menu);
			this.current_item_menu_obj._onclick = this.current_item_menu_obj.onclick;
			this.current_item_menu_obj.onmouseover = this.current_item_menu_obj.onmouseout = this.current_item_menu_obj.onclick = this.current_menu_mouseover;
			ipsmenu.menu_registered[this.current_item_menu]._open = ipsmenu.menu_registered[this.current_item_menu].open;
			ipsmenu.menu_registered[this.current_item_menu].open = this.current_menu_button_show;
		}
	};
	this.init_current_menu = function (obj) {
		var menu = '';
		var ipsatt = (obj.add_app_id) ? ipsattach[obj.add_app_id] : ipsattach;
		try {
			menu = document.getElementById(ipsatt.current_item_menu + '_menu');
			menu.innerHTML = '';
		} catch (error) {
			menu = document.createElement('div');
		}
		var _items = 0;
		menu.add_app_id = (obj.add_app_id) ? obj.add_app_id : '';
		menu.id = ipsatt.current_item_menu + '_menu';
		menu.className = 'attach-popupmenu';
		menu.style.display = 'none';
		menu.style.cursor = 'default';
		menu.style.padding = '3px';
		menu.style.width = 'auto';
		menu.style.overflow = 'auto';
		for (var i in ipsatt.current_items) {
			var option = document.createElement('div');
			option.add_app_id = (obj.add_app_id) ? obj.add_app_id : '';
			option.innerHTML = ipsatt.current_items[i];
			option.style.width = ipsatt.current_item_menu_obj.style.width;
			option.className = 'attach-menu-row';
			option.onmouseover = option.onmouseout = ipsatt.menu_onmouse_event;
			menu.appendChild(option);
			_items++;
		}
		if (!_items) {
			var option = document.createElement('div');
			option.add_app_id = (obj.add_app_id) ? obj.add_app_id : '';
			option.innerHTML = ipsatt.lang['no_items'];
			option.style.width = ipsatt.current_item_menu_obj.style.width;
			option.className = 'attach-menu-row';
			option.onmouseover = option.onmouseout = ipsatt.menu_onmouse_event;
			menu.appendChild(option);
		}
		ipsatt.set_attach_count_status(obj, _items);
		ipsatt.iframe_parent_div_obj.appendChild(menu);
		ipsclass.set_unselectable(menu);
	};
	this.set_attach_count_status = function (obj, _items) {
		var ipsatt = (obj.add_app_id) ? ipsattach[this.add_app_id] : ipsattach;
		if (!ipsatt.show_attach_menu) {
			return;
		}
		if (typeof (_items) == 'undefined') {
			_items = 0;
			for (var i in ipsatt.current_items) {
				_items++;
			}
		}
		try {
			ipsatt.dropdown_text_obj.innerHTML = " (" + _items + ")";
		} catch (error) {}
	};
	this.inArray=function(t,a){
		if(!(typeof(a)=='object'&&(a instanceof Array)))return false;
		for(var i=0;i<a.length;a++)
			if(a[i]===t)return true;
		return false;
	};
	this.add_current_item = function (attach_id, attach_name, attach_size, attach_image) {
		var vipsatt = '';
		var ipsatt = null;
		if (this.add_app_id) {
			ipsatt = ipsattach[this.add_app_id];
			vipsatt = "ipsattach[" + this.add_app_id + "]";
		} else {
			ipsatt = ipsattach;
			vipsatt = "ipsattach";
		}
		if (attach_name.length > 22) {
			var _ext = attach_name.replace(/^.*\.(\S+?)$/, "$1");
			var _main = attach_name.replace(/^(.*)\.\S+?$/, "$1");
			var _a = _main.substr(0, 8);
			var _b = _main.substr(_main.length - 8, _main.length);
			attach_name = _a + '...' + _b + '.' + _ext;
		}
		var html = '';
		var add_code_js = '';
		if (ipsatt.can_add_filecode) {
			add_code_js = ' onclick="' + vipsatt + '.add_attachment_into_editor(' + attach_id + ', \'' + attach_name + '\')"';
			html += '<img onclick="' + vipsatt + '.add_attachment_into_editor(' + attach_id + ', \'' + attach_name + '\')" src="' + this.images_url + '/folder_attach_images/attach_add.png" title="' + this.lang['attach_insert'] + '" style="vertical-align:middle" border="0" alt="" /> &nbsp;';
		}
		html += '<img onclick="' + vipsatt + '.remove_attachment(' + attach_id + ', \'' + this.add_app_id + '\');" src="' + this.images_url + '/folder_attach_images/attach_remove.png" title="' + this.lang['attach_remove'] + ': [' + attach_id + ']' + '" style="vertical-align:middle" border="0" alt="" />&nbsp;';
		html += '<img src="' + this.images_url + '/' + attach_image + '" style="vertical-align:middle" border="0" alt="" /> &nbsp;';
		html += '<span style="font-weight:bold"' + add_code_js + '>' + attach_name + '</span>';
		html += ' <span class="desctext">&nbsp;(' + attach_size + ')</span>';
		//this.current_items[this.current_items.length] = html;
		this.current_items[attach_id]=html;
		var t='';
		for(var i in this.current_items)t+=i+',';
		this.frm_filelist.value=t.replace(/(^,|,$)/g,'');
	};
	this.add_attachment_into_editor = function (id, name) {
		var tag = "[attachment=" + id + ":" + name.replace(/\[|\]/g,'_') + "]";
		if(this.edid)
			ipsclass.add_editor_contents(tag,this.edid);
		else
			ipsclass.add_editor_contents(tag);
	};
	this.remove_attachment = function (id, add_app_id) {
		var ipsatt = (add_app_id) ? ipsattach[add_app_id] : ipsattach;
		menu_action_close();
		if (confirm(ipsatt.lang['remove_warn'] + ' [' + id + ']')) {
			ipsatt.iframe_obj.src = ipsatt.iframe_remove_url + "&attach_id=" + id;
			ipsatt.iframe_on_un_load('attach_removal');
			var na={};
			var t='';
			for(var i in ipsatt.current_items){if(i!=id){na[i]=ipsatt.current_items[i];t+=i+',';}}
			ipsatt.current_items=na;
			ipsatt.frm_filelist.value=t.replace(/(^,|,$)/g,'');
		}
	};
	this.show_attach_box = function () {
		var iheight = parseInt(this.iframe_add_to_div_obj.style.height);
		var iwidth = this.iframe_add_to_div_obj.style.width;
		if (!this.iframe_obj) {
			this.iframe_obj = document.createElement('IFRAME');
			this.iframe_obj.add_app_id = (this.add_app_id) ? this.add_app_id : '';
			this.iframe_obj.src = this.iframe_init_url;
			this.iframe_obj.id = this.iframe_id;
			this.iframe_obj.name = this.iframe_id;
			this.iframe_obj.scrolling = 'no';
			this.iframe_obj.frameBorder = 'no';
			this.iframe_obj.border = '0';
			this.iframe_obj.className = this.iframe_classname;
			this.iframe_obj.style.width = iwidth ? iwidth : '100%';
			this.iframe_obj.style.height = iheight ? iheight - 5 + 'px' : this.iframe_height + 'px';
			this.iframe_obj.style.overflow = 'hidden';
			this.iframe_obj.style.display = '';
			this.iframe_obj.msg_open = 1;
			this.iframe_obj.iframe_loaded = 0;
			this.iframe_obj.iframe_init = 0;
			this.iframe_add_to_div_obj.appendChild(this.iframe_obj);
			if (is_ie) {
				this.iframe_obj.style.backgroundColor = 'transparent';
				this.iframe_obj.allowTransparency = true;
				this.iframe_obj.onreadystatechange = this.iframe_on_load_ie;
			} else {
				this.iframe_obj.onload = this.iframe_onload;
			}
			if (is_safari) {
				this.iframe_parent_div_obj.style.display = '';
			} else {
				this.iframe_parent_div_obj.style.display = 'none';
			}
		}
		if (typeof this.message_div_obj != 'undefined') {
			this.message_div_obj.style.width = typeof this.iframe_parent_div_obj.style.width != 'undefined' ? this.iframe_parent_div_obj.style.width : '100%';
			this.message_div_obj.style.height = this.iframe_parent_div_obj.offsetHeight ? this.iframe_parent_div_obj.offsetHeight : '120px';
			this.message_div_obj.style.position = 'relative';
			this.message_div_obj.style.display = 'none';
		}
		this.show_message(this.lang['init_progress']);
	};
	this.init_status_bar = function () {
		var _html = '';
		var ipsatt = null;
		if (this.add_app_id) {
			ipsatt = ipsattach[this.add_app_id];
		} else {
			ipsatt = ipsattach;
		}
		if (typeof ipsatt.status_div_obj != 'undefined') {
			try {
				ipsatt.status_msg = ipsatt.lang[ipsatt.status_msg];
			} catch (error) {}
			if (ipsatt.status_msg) {
				if (ipsatt.status_is_error) {
					_html += "<img src='" + this.images_url + "/folder_attach_images/attach_error.png' border='0' style='vertical-align:middle' alt='" + ipb_global_lang['general_error'] + "' />&nbsp;" + ipsatt.status_msg;
				} else {
					_html += "<img src='" + this.images_url + "/folder_attach_images/attach_ok.png' border='0' style='vertical-align:middle' alt='" + ipb_global_lang['general_OK'] + "' />&nbsp;" + ipsatt.status_msg;
				}
			}
			ipsatt.status_div_obj.innerHTML = _html;
		}
		ipsatt.has_upload_pending = 0;
	};
	this.iframe_on_un_load = function (msg) {
		var ipsatt = (this.add_app_id) ? ipsattach[this.add_app_id] : ipsattach;
		msg = typeof (msg) != 'undefined' ? ipsatt.lang[msg] : ipsatt.lang['uploading_file'];
		if (this.iframe_obj.iframe_loaded) {
			this.iframe_obj.iframe_loaded = 0;
			this.iframe_obj.msg_open = 1;
			ipsatt.show_message(msg);
		}
	};
	this.iframe_onload = function (e) {
		var evt = window.event || e;
		if (!evt.target) evt.target = evt.srcElement;
		var ipsatt = (evt.target.add_app_id) ? ipsattach[evt.target.add_app_id] : ipsattach;
		if (!this.iframe_init) {
			this.iframe_init = 1;
			this.iframe_loaded = 1;
			if (this.msg_open) {
				this.msg_open = 0;
				ipsatt.hide_message();
			}
			window.frames[this.id].document.onmousedown = menu_action_close;
		} else {
			this.iframe_loaded = 1;
			window.frames[this.id].document.onmousedown = menu_action_close;
			if (this.msg_open) {
				this.msg_open = 0;
				ipsatt.hide_message();
			}
			if (ipsatt.onload_function) {
				ipsatt.onload_function();
			}
		}
		ipsatt.init_status_bar();
		ipsatt.set_attach_count_status(evt.target);
	};
	this.iframe_on_load_ie = function (e) {
		var evt = window.event || e;
		if (!evt.target) evt.target = evt.srcElement;
		var ipsatt = (evt.target.add_app_id) ? ipsattach[evt.target.add_app_id] : ipsattach;
		if (this.readyState == 'complete') {
			if (!this.iframe_init) {
				this.iframe_init = 1;
				this.iframe_loaded = 1;
				window.frames[this.id].document.onmousedown = menu_action_close;
				if (this.msg_open) {
					this.msg_open = 0;
					ipsatt.hide_message();
				}
			} else {
				this.iframe_loaded = 1;
				window.frames[this.id].document.onmousedown = menu_action_close;
				if (this.msg_open) {
					this.msg_open = 0;
					ipsatt.hide_message();
				}
			}
			window.frames[this.id].document.getElementsByTagName('html')[0].style.background = 'transparent';
			if (ipsatt.onload_function) {
				ipsatt.onload_function();
			}
			ipsatt.init_status_bar();
			ipsatt.set_attach_count_status(evt.target);
		}
	};
	this.show_message = function (msg) {
		var ipsatt = (this.add_app_id) ? ipsattach[this.add_app_id] : ipsattach;
		if (is_safari) {
			return;
		}
		if (typeof ipsatt.message_div_obj != 'undefined') {
			ipsatt.message_div_obj.style.display = '';
			ipsatt.message_span_obj.innerHTML = msg;
			ipsatt.iframe_parent_div_obj.style.display = 'none';
		}
	};
	this.hide_message = function () {
		var ipsatt = (this.add_app_id) ? ipsattach[this.add_app_id] : ipsattach;
		if (is_safari) {
			return;
		}
		if (typeof ipsatt.message_div_obj != 'undefined') {
			ipsatt.message_div_obj.style.display = 'none';
			ipsatt.message_span_obj.innerHTML = '';
			ipsatt.iframe_parent_div_obj.style.display = '';
		}
	};
	this.current_menu_mouseover = function (e) {
		e = ipsclass.cancel_bubble(e, true);
		if (e.type == 'click'&&this._onclick) {
			this._onclick(e);
		}
	};
	this.current_menu_button_show = function (obj) {
		var ipsatt = (obj.add_app_id) ? ipsattach[obj.add_app_id] : ipsattach;
		ipsatt.init_current_menu(obj);
		this._open(obj);
	};
	this.menu_onmouse_event = function (e) {
		var evt = window.event || e;
		if (!evt.target) evt.target = evt.srcElement;
		var ipsatt = (evt.target.add_app_id) ? ipsattach[evt.target.add_app_id] : ipsattach_css;
		e = ipsclass.cancel_bubble(e, true);
		if (e.type == 'mouseover') {
			this.className = ipsatt.css_menu_row_over;
		} else if (e.type == 'mouseout') {
			this.className = ipsatt.css_menu_row;
		}
	};
}
