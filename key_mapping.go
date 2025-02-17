package activitylogger

var KeyMap = map[int]string{
	0:   "KEY_RESERVED",
	1:   "KEY_ESC",
	2:   "KEY_1",
	3:   "KEY_2",
	4:   "KEY_3",
	5:   "KEY_4",
	6:   "KEY_5",
	7:   "KEY_6",
	8:   "KEY_7",
	9:   "KEY_8",
	10:  "KEY_9",
	11:  "KEY_0",
	12:  "KEY_MINUS",
	13:  "KEY_EQUAL",
	14:  "KEY_BACKSPACE",
	15:  "KEY_TAB",
	16:  "KEY_Q",
	17:  "KEY_W",
	18:  "KEY_E",
	19:  "KEY_R",
	20:  "KEY_T",
	21:  "KEY_Y",
	22:  "KEY_U",
	23:  "KEY_I",
	24:  "KEY_O",
	25:  "KEY_P",
	26:  "KEY_LEFTBRACE",
	27:  "KEY_RIGHTBRACE",
	28:  "KEY_ENTER",
	29:  "KEY_LEFTCTRL",
	30:  "KEY_A",
	31:  "KEY_S",
	32:  "KEY_D",
	33:  "KEY_F",
	34:  "KEY_G",
	35:  "KEY_H",
	36:  "KEY_J",
	37:  "KEY_K",
	38:  "KEY_L",
	39:  "KEY_SEMICOLON",
	40:  "KEY_APOSTROPHE",
	41:  "KEY_GRAVE",
	42:  "KEY_LEFTSHIFT",
	43:  "KEY_BACKSLASH",
	44:  "KEY_Z",
	45:  "KEY_X",
	46:  "KEY_C",
	47:  "KEY_V",
	48:  "KEY_B",
	49:  "KEY_N",
	50:  "KEY_M",
	51:  "KEY_COMMA",
	52:  "KEY_DOT",
	53:  "KEY_SLASH",
	54:  "KEY_RIGHTSHIFT",
	55:  "KEY_KPASTERISK",
	56:  "KEY_LEFTALT",
	57:  "KEY_SPACE",
	58:  "KEY_CAPSLOCK",
	59:  "KEY_F1",
	60:  "KEY_F2",
	61:  "KEY_F3",
	62:  "KEY_F4",
	63:  "KEY_F5",
	64:  "KEY_F6",
	65:  "KEY_F7",
	66:  "KEY_F8",
	67:  "KEY_F9",
	68:  "KEY_F10",
	69:  "KEY_NUMLOCK",
	70:  "KEY_SCROLLLOCK",
	71:  "KEY_KP7",
	72:  "KEY_KP8",
	73:  "KEY_KP9",
	74:  "KEY_KPMINUS",
	75:  "KEY_KP4",
	76:  "KEY_KP5",
	77:  "KEY_KP6",
	78:  "KEY_KPPLUS",
	79:  "KEY_KP1",
	80:  "KEY_KP2",
	81:  "KEY_KP3",
	82:  "KEY_KP0",
	83:  "KEY_KPDOT",
	85:  "KEY_ZENKAKUHANKAKU",
	86:  "KEY_102ND",
	87:  "KEY_F11",
	88:  "KEY_F12",
	89:  "KEY_RO",
	90:  "KEY_KATAKANA",
	91:  "KEY_HIRAGANA",
	92:  "KEY_HENKAN",
	93:  "KEY_KATAKANAHIRAGANA",
	94:  "KEY_MUHENKAN",
	95:  "KEY_KPJPCOMMA",
	96:  "KEY_KPENTER",
	97:  "KEY_RIGHTCTRL",
	98:  "KEY_KPSLASH",
	99:  "KEY_SYSRQ",
	100: "KEY_RIGHTALT",
	101: "KEY_LINEFEED",
	102: "KEY_HOME",
	103: "KEY_UP",
	104: "KEY_PAGEUP",
	105: "KEY_LEFT",
	106: "KEY_RIGHT",
	107: "KEY_END",
	108: "KEY_DOWN",
	109: "KEY_PAGEDOWN",
	110: "KEY_INSERT",
	111: "KEY_DELETE",
	112: "KEY_MACRO",
	113: "KEY_MUTE",
	114: "KEY_VOLUMEDOWN",
	115: "KEY_VOLUMEUP",
	116: "KEY_POWER",
	117: "KEY_KPEQUAL",
	118: "KEY_KPPLUSMINUS",
	119: "KEY_PAUSE",
	120: "KEY_SCALE",
	121: "KEY_KPCOMMA",
	122: "KEY_HANGEUL",
	123: "KEY_HANJA",
	124: "KEY_YEN",
	125: "KEY_LEFTMETA",
	126: "KEY_RIGHTMETA",
	127: "KEY_COMPOSE",
	128: "KEY_STOP",
	129: "KEY_AGAIN",
	130: "KEY_PROPS",
	131: "KEY_UNDO",
	132: "KEY_FRONT",
	133: "KEY_COPY",
	134: "KEY_OPEN",
	135: "KEY_PASTE",
	136: "KEY_FIND",
	137: "KEY_CUT",
	138: "KEY_HELP",
	139: "KEY_MENU",
	140: "KEY_CALC",
	141: "KEY_SETUP",
	142: "KEY_SLEEP",
	143: "KEY_WAKEUP",
	144: "KEY_FILE",
	145: "KEY_SENDFILE",
	146: "KEY_DELETEFILE",
	147: "KEY_XFER",
	148: "KEY_PROG1",
	149: "KEY_PROG2",
	150: "KEY_WWW",
	151: "KEY_MSDOS",
	152: "KEY_COFFEE",
	154: "KEY_CYCLEWINDOWS",
	155: "KEY_MAIL",
	156: "KEY_BOOKMARKS",
	157: "KEY_COMPUTER",
	158: "KEY_BACK",
	159: "KEY_FORWARD",
	160: "KEY_CLOSECD",
	161: "KEY_EJECTCD",
	162: "KEY_EJECTCLOSECD",
	163: "KEY_NEXTSONG",
	164: "KEY_PLAYPAUSE",
	165: "KEY_PREVIOUSSONG",
	166: "KEY_STOPCD",
	167: "KEY_RECORD",
	168: "KEY_REWIND",
	169: "KEY_PHONE",
	170: "KEY_ISO",
	171: "KEY_CONFIG",
	172: "KEY_HOMEPAGE",
	173: "KEY_REFRESH",
	174: "KEY_EXIT",
	175: "KEY_MOVE",
	176: "KEY_EDIT",
	177: "KEY_SCROLLUP",
	178: "KEY_SCROLLDOWN",
	179: "KEY_KPLEFTPAREN",
	180: "KEY_KPRIGHTPAREN",
	181: "KEY_NEW",
	182: "KEY_REDO",
	183: "KEY_F13",
	184: "KEY_F14",
	185: "KEY_F15",
	186: "KEY_F16",
	187: "KEY_F17",
	188: "KEY_F18",
	189: "KEY_F19",
	190: "KEY_F20",
	191: "KEY_F21",
	192: "KEY_F22",
	193: "KEY_F23",
	194: "KEY_F24",
	200: "KEY_PLAYCD",
	201: "KEY_PAUSECD",
	202: "KEY_PROG3",
	203: "KEY_PROG4",
	205: "KEY_SUSPEND",
	206: "KEY_CLOSE",
	207: "KEY_PLAY",
	208: "KEY_FASTFORWARD",
	209: "KEY_BASSBOOST",
	210: "KEY_PRINT",
	211: "KEY_HP",
	212: "KEY_CAMERA",
	213: "KEY_SOUND",
	214: "KEY_QUESTION",
	215: "KEY_EMAIL",
	216: "KEY_CHAT",
	217: "KEY_SEARCH",
	218: "KEY_CONNECT",
	219: "KEY_FINANCE",
	220: "KEY_SPORT",
	221: "KEY_SHOP",
	222: "KEY_ALTERASE",
	223: "KEY_CANCEL",
	224: "KEY_BRIGHTNESSDOWN",
	225: "KEY_BRIGHTNESSUP",
	226: "KEY_MEDIA",
	227: "KEY_SWITCHVIDEOMODE",
	228: "KEY_KBDILLUMTOGGLE",
	229: "KEY_KBDILLUMDOWN",
	230: "KEY_KBDILLUMUP",
	231: "KEY_SEND",
	232: "KEY_REPLY",
	233: "KEY_FORWARDMAIL",
	234: "KEY_SAVE",
	235: "KEY_DOCUMENTS",
	236: "KEY_BATTERY",
	237: "KEY_BLUETOOTH",
	238: "KEY_WLAN",
	239: "KEY_UWB",
	240: "KEY_UNKNOWN",
	246: "KEY_WWAN",
	247: "KEY_RFKILL",
	248: "KEY_MICMUTE",
	272: "BTN_LEFT",
	273: "BTN_RIGHT",
	274: "BTN_MIDDLE",
	275: "BTN_SIDE",
	276: "BTN_EXTRA",
	277: "BTN_FORWARD",
	278: "BTN_BACK",
	279: "BTN_TASK",
	352: "KEY_OK",
	353: "KEY_SELECT",
	354: "KEY_GOTO",
	355: "KEY_CLEAR",
	356: "KEY_POWER2",
	357: "KEY_OPTION",
	358: "KEY_INFO",
	359: "KEY_TIME",
	360: "KEY_VENDOR",
	361: "KEY_ARCHIVE",
	362: "KEY_PROGRAM",
	363: "KEY_CHANNEL",
	364: "KEY_FAVORITES",
	365: "KEY_EPG",
	366: "KEY_PVR",
	367: "KEY_MHP",
	368: "KEY_LANGUAGE",
	369: "KEY_TITLE",
	370: "KEY_SUBTITLE",
	371: "KEY_ANGLE",
	373: "KEY_MODE",
	374: "KEY_KEYBOARD",
	376: "KEY_PC",
	377: "KEY_TV",
	378: "KEY_TV2",
	379: "KEY_VCR",
	380: "KEY_VCR2",
	381: "KEY_SAT",
	382: "KEY_SAT2",
	383: "KEY_CD",
	384: "KEY_TAPE",
	385: "KEY_RADIO",
	386: "KEY_TUNER",
	387: "KEY_PLAYER",
	388: "KEY_TEXT",
	389: "KEY_DVD",
	390: "KEY_AUX",
	391: "KEY_MP3",
	392: "KEY_AUDIO",
	393: "KEY_VIDEO",
	394: "KEY_DIRECTORY",
	395: "KEY_LIST",
	396: "KEY_MEMO",
	397: "KEY_CALENDAR",
	398: "KEY_RED",
	399: "KEY_GREEN",
	400: "KEY_YELLOW",
	401: "KEY_BLUE",
	402: "KEY_CHANNELUP",
	403: "KEY_CHANNELDOWN",
	404: "KEY_FIRST",
	405: "KEY_LAST",
	406: "KEY_AB",
	407: "KEY_NEXT",
	408: "KEY_RESTART",
	409: "KEY_SLOW",
	410: "KEY_SHUFFLE",
	411: "KEY_BREAK",
	412: "KEY_PREVIOUS",
	413: "KEY_DIGITS",
	414: "KEY_TEEN",
	415: "KEY_TWEN",
	416: "KEY_VIDEOPHONE",
	417: "KEY_GAMES",
	418: "KEY_ZOOMIN",
	419: "KEY_ZOOMOUT",
	420: "KEY_ZOOMRESET",
	421: "KEY_WORDPROCESSOR",
	422: "KEY_EDITOR",
	423: "KEY_SPREADSHEET",
	424: "KEY_GRAPHICSEDITOR",
	425: "KEY_PRESENTATION",
	426: "KEY_DATABASE",
	427: "KEY_NEWS",
	428: "KEY_VOICEMAIL",
	429: "KEY_ADDRESSBOOK",
	430: "KEY_MESSENGER",
	431: "KEY_DISPLAYTOGGLE",
	432: "KEY_SPELLCHECK",
	433: "KEY_LOGOFF",
	434: "KEY_DOLLAR",
	435: "KEY_EURO",
	436: "KEY_FRAMEBACK",
	437: "KEY_FRAMEFORWARD",
	440: "KEY_10CHANNELSUP",
	441: "KEY_10CHANNELSDOWN",
	442: "KEY_IMAGES",
	464: "KEY_FN",
	576: "KEY_BUTTONCONFIG",
	577: "KEY_TASKMANAGER",
	578: "KEY_JOURNAL",
	579: "KEY_CONTROLPANEL",
	580: "KEY_APPSELECT",
	581: "KEY_SCREENSAVER",
	582: "KEY_VOICECOMMAND",
	583: "KEY_ASSISTANT",
	586: "KEY_DICTATE",
	590: "KEY_ACCESSIBILITY",
	627: "KEY_VOD",
	628: "KEY_UNMUTE",
	629: "KEY_FASTREVERSE",
	630: "KEY_SLOWREVERSE",
	631: "KEY_DATA",
	639: "KEY_SOS",
	656: "KEY_MACRO1",
	657: "KEY_MACRO2",
	658: "KEY_MACRO3",
	659: "KEY_MACRO4",
	660: "KEY_MACRO5",
	661: "KEY_MACRO6",
	662: "KEY_MACRO7",
	663: "KEY_MACRO8",
	664: "KEY_MACRO9",
	665: "KEY_MACRO10",
	666: "KEY_MACRO11",
	667: "KEY_MACRO12",
	668: "KEY_MACRO13",
	669: "KEY_MACRO14",
	670: "KEY_MACRO15",
	671: "KEY_MACRO16",
	672: "KEY_MACRO17",
	673: "KEY_MACRO18",
	674: "KEY_MACRO19",
	675: "KEY_MACRO20",
	676: "KEY_MACRO21",
	677: "KEY_MACRO22",
	678: "KEY_MACRO23",
	679: "KEY_MACRO24",
	680: "KEY_MACRO25",
	681: "KEY_MACRO26",
	682: "KEY_MACRO27",
	683: "KEY_MACRO28",
	684: "KEY_MACRO29",
	685: "KEY_MACRO30",
	767: "KEY_MAX",
}
