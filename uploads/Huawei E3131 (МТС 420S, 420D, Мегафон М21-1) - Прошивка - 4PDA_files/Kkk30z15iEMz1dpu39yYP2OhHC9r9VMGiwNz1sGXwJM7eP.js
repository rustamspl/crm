//============================================
// Global Announcements mod for IPB 2.1 - 2.3
// Version 2.6
//
// (c) 2005-2007, DINI
//
//--------------------------------------------
//
// JAVA SCRIPT FUNCTIONS
//
//============================================

function toggleglobalmess( fid, add, update, tm ) {
	saved = new Array();
	clean = new Array();

	if ( tmp = my_getcookie('globalmesscollapse') )
	{
		saved = tmp.split(",");
	}

	for( i = 0 ; i < saved.length; i++ )
	{
		if ( saved[i] != fid && saved[i] != "" )
		{
			clean[clean.length] = saved[i];
		}
	}

	if ( add )
	{
		my_show_div( my_getbyid( 'gc_'+fid  ) );
		my_hide_div( my_getbyid( 'go_'+fid  ) );
	}
	else
	{
		my_show_div( my_getbyid( 'go_'+fid  ) );
		my_hide_div( my_getbyid( 'gc_'+fid  ) );
	}

	my_setcookie( 'globalmesscollapse', '1'+add, 1 );

	if( update)
	{
		my_setcookie( 'globalmessupdmess', update, 1 );
		my_setcookie( 'globalmessupdtime', tm || (new Date()).getTime()/1000, 1 );
	}
}