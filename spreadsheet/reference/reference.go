//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package reference ;import (_cdd "errors";_a "fmt";_ab "github.com/unidoc/unioffice/spreadsheet/update";_ceb "regexp";_da "strconv";_ac "strings";);

// Update updates reference to point one of the neighboring cells with respect to the update type after removing a row/column.
func (_d *CellReference )Update (updateType _ab .UpdateAction )*CellReference {switch updateType {case _ab .UpdateActionRemoveColumn :_gf :=_d ;_gf .ColumnIdx =_d .ColumnIdx -1;_gf .Column =IndexToColumn (_gf .ColumnIdx );return _gf ;default:return _d ;};};

// ParseColumnReference parses a column reference of the form 'Sheet1!A' and splits it
// into sheet name and column segments.
func ParseColumnReference (s string )(ColumnReference ,error ){s =_ac .TrimSpace (s );if len (s )< 1{return ColumnReference {},_cdd .New ("\u0063\u006f\u006c\u0075\u006d\u006e \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0063\u0068a\u0072a\u0063\u0074\u0065\u0072");};_gb :=ColumnReference {};_eaa :=_ac .Split (s ,"\u0021");if len (_eaa )==2{_gb .SheetName =_eaa [0];s =_eaa [1];};if s [0]=='$'{_gb .AbsoluteColumn =true ;s =s [1:];};if !_gff .MatchString (s ){return ColumnReference {},_cdd .New ("\u0063\u006f\u006c\u0075\u006dn\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0041\u0020\u0061\u006e\u0064\u0020\u005a\u005a");};_gb .Column =s ;_gb .ColumnIdx =ColumnToIndex (_gb .Column );return _gb ,nil ;};

// ColumnReference is a parsed reference to a column.  Input is of the form 'A',
// '$C', etc.
type ColumnReference struct{ColumnIdx uint32 ;Column string ;AbsoluteColumn bool ;SheetName string ;};var _gff =_ceb .MustCompile ("^\u005b\u0061\u002d\u007aA-\u005a]\u0028\u005b\u0061\u002d\u007aA\u002d\u005a\u005d\u003f\u0029\u0024");

// ColumnToIndex maps a column to a zero based index (e.g. A = 0, B = 1, AA = 26)
func ColumnToIndex (col string )uint32 {col =_ac .ToUpper (col );_eef :=uint32 (0);for _ ,_g :=range col {_eef *=26;_eef +=uint32 (_g -'A'+1);};return _eef -1;};

// ParseColumnRangeReference splits a range reference of the form "A:B" into its
// components.
func ParseColumnRangeReference (s string )(_cdf ,_e ColumnReference ,_cea error ){_be :="";_abb :=_ac .Split (s ,"\u0021");if len (_abb )==2{_be =_abb [0];s =_abb [1];};_ed :=_ac .Split (s ,"\u003a");if len (_ed )!=2{return ColumnReference {},ColumnReference {},_cdd .New ("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074");};if _be !=""{_ed [0]=_be +"\u0021"+_ed [0];_ed [1]=_be +"\u0021"+_ed [1];};_fd ,_cea :=ParseColumnReference (_ed [0]);if _cea !=nil {return ColumnReference {},ColumnReference {},_cea ;};_de ,_cea :=ParseColumnReference (_ed [1]);if _cea !=nil {return ColumnReference {},ColumnReference {},_cea ;};return _fd ,_de ,nil ;};

// IndexToColumn maps a column number to a column name (e.g. 0 = A, 1 = B, 26 = AA)
func IndexToColumn (col uint32 )string {var _fb [64+1]byte ;_ae :=len (_fb );_ee :=col ;const _aa =26;for _ee >=_aa {_ae --;_b :=_ee /_aa ;_fb [_ae ]=byte ('A'+uint (_ee -_b *_aa ));_ee =_b -1;};_ae --;_fb [_ae ]=byte ('A'+uint (_ee ));return string (_fb [_ae :]);};

// Update updates reference to point one of the neighboring columns with respect to the update type after removing a row/column.
func (_beb *ColumnReference )Update (updateType _ab .UpdateAction )*ColumnReference {switch updateType {case _ab .UpdateActionRemoveColumn :_c :=_beb ;_c .ColumnIdx =_beb .ColumnIdx -1;_c .Column =IndexToColumn (_c .ColumnIdx );return _c ;default:return _beb ;};};

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference (s string )(CellReference ,error ){s =_ac .TrimSpace (s );if len (s )< 2{return CellReference {},_cdd .New ("\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066e\u0072\u0065\u006ece\u0020\u006d\u0075\u0073\u0074\u0020h\u0061\u0076\u0065\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077o\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0073");};_ge :=CellReference {};_ba :=_ac .Split (s ,"\u0021");if len (_ba )==2{_ge .SheetName =_ba [0];s =_ba [1];};if s [0]=='$'{_ge .AbsoluteColumn =true ;s =s [1:];};_f :=-1;_abg :for _cd :=0;_cd < len (s );_cd ++{switch {case s [_cd ]>='0'&&s [_cd ]<='9'||s [_cd ]=='$':_f =_cd ;break _abg ;};};switch _f {case 0:return CellReference {},_a .Errorf ("\u006e\u006f\u0020\u006cet\u0074\u0065\u0072\u0020\u0070\u0072\u0065\u0066\u0069\u0078\u0020\u0069\u006e\u0020%\u0073",s );case -1:return CellReference {},_a .Errorf ("\u006eo\u0020d\u0069\u0067\u0069\u0074\u0073\u0020\u0069\u006e\u0020\u0025\u0073",s );};_ge .Column =s [0:_f ];if s [_f ]=='$'{_ge .AbsoluteRow =true ;_f ++;};_ge .ColumnIdx =ColumnToIndex (_ge .Column );_dc ,_gd :=_da .ParseUint (s [_f :],10,32);if _gd !=nil {return CellReference {},_a .Errorf ("e\u0072\u0072\u006f\u0072 p\u0061r\u0073\u0069\u006e\u0067\u0020r\u006f\u0077\u003a\u0020\u0025\u0073",_gd );};_ge .RowIdx =uint32 (_dc );return _ge ,nil ;};

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference (s string )(_bb ,_fc CellReference ,_ff error ){_df :="";_ddb :=_ac .Split (s ,"\u0021");if len (_ddb )==2{_df =_ddb [0];s =_ddb [1];};_ca :=_ac .Split (s ,"\u003a");if len (_ca )!=2{return CellReference {},CellReference {},_cdd .New ("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074");};if _df !=""{_ca [0]=_df +"\u0021"+_ca [0];_ca [1]=_df +"\u0021"+_ca [1];};_bd ,_ff :=ParseCellReference (_ca [0]);if _ff !=nil {return CellReference {},CellReference {},_ff ;};_bg ,_ff :=ParseCellReference (_ca [1]);if _ff !=nil {return CellReference {},CellReference {},_ff ;};return _bd ,_bg ,nil ;};

// String returns a string representation of CellReference.
func (_ea CellReference )String ()string {_dd :=make ([]byte ,0,4);if _ea .AbsoluteColumn {_dd =append (_dd ,'$');};_dd =append (_dd ,_ea .Column ...);if _ea .AbsoluteRow {_dd =append (_dd ,'$');};_dd =_da .AppendInt (_dd ,int64 (_ea .RowIdx ),10);return string (_dd );};

// String returns a string representation of ColumnReference.
func (_geb ColumnReference )String ()string {_deg :=make ([]byte ,0,4);if _geb .AbsoluteColumn {_deg =append (_deg ,'$');};_deg =append (_deg ,_geb .Column ...);return string (_deg );};

// CellReference is a parsed reference to a cell.  Input is of the form 'A1',
// '$C$2', etc.
type CellReference struct{RowIdx uint32 ;ColumnIdx uint32 ;Column string ;AbsoluteColumn bool ;AbsoluteRow bool ;SheetName string ;};