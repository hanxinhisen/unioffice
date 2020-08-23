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

package sharedTypes ;import (_cfb "encoding/xml";_dfg "fmt";_cd "regexp";);var ST_UniversalMeasurePatternRe =_cd .MustCompile (ST_UniversalMeasurePattern );func (_dbg ST_AlgType )ValidateWithPath (path string )error {switch _dbg {case 0,1,2:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_dbg ));};return nil ;};func (_ace ST_XAlign )String ()string {switch _ace {case 0:return "";case 1:return "\u006c\u0065\u0066\u0074";case 2:return "\u0063\u0065\u006e\u0074\u0065\u0072";case 3:return "\u0072\u0069\u0067h\u0074";case 4:return "\u0069\u006e\u0073\u0069\u0064\u0065";case 5:return "\u006fu\u0074\u0073\u0069\u0064\u0065";};return "";};func (_afc ST_YAlign )ValidateWithPath (path string )error {switch _afc {case 0,1,2,3,4,5,6:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_afc ));};return nil ;};func (_bab ST_CalendarType )Validate ()error {return _bab .ValidateWithPath ("")};const (ST_TrueFalseUnset ST_TrueFalse =0;ST_TrueFalseT ST_TrueFalse =1;ST_TrueFalseF ST_TrueFalse =2;ST_TrueFalseTrue ST_TrueFalse =3;ST_TrueFalseFalse ST_TrueFalse =4;);const (ST_ConformanceClassUnset ST_ConformanceClass =0;ST_ConformanceClassStrict ST_ConformanceClass =1;ST_ConformanceClassTransitional ST_ConformanceClass =2;);var ST_PositiveFixedPercentagePatternRe =_cd .MustCompile (ST_PositiveFixedPercentagePattern );type ST_AlgType byte ;const ST_PositiveFixedPercentagePattern ="\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039\u005d\u003f\u0029\u0029\u0028\u005c\u002e[\u0030\u002d\u0039\u005d\u005b0\u002d\u0039]\u003f\u0029\u003f\u0025";func (_fgg ST_AlgClass )Validate ()error {return _fgg .ValidateWithPath ("")};type ST_TrueFalseBlank byte ;const ST_PositivePercentagePattern ="\u005b0\u002d9\u005d\u002b\u0028\u005c\u002e[\u0030\u002d9\u005d\u002b\u0029\u003f\u0025";const (ST_XAlignUnset ST_XAlign =0;ST_XAlignLeft ST_XAlign =1;ST_XAlignCenter ST_XAlign =2;ST_XAlignRight ST_XAlign =3;ST_XAlignInside ST_XAlign =4;ST_XAlignOutside ST_XAlign =5;);func (_f *ST_AlgClass )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_f =0;case "\u0068\u0061\u0073\u0068":*_f =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_f =2;};return nil ;};func (_gea *ST_VerticalAlignRun )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_dce ,_dab :=d .Token ();if _dab !=nil {return _dab ;};if _e ,_cgc :=_dce .(_cfb .EndElement );_cgc &&_e .Name ==start .Name {*_gea =1;return nil ;};if _bf ,_adff :=_dce .(_cfb .CharData );!_adff {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_dce );}else {switch string (_bf ){case "":*_gea =0;case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":*_gea =1;case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":*_gea =2;case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":*_gea =3;};};_dce ,_dab =d .Token ();if _dab !=nil {return _dab ;};if _daed ,_gbd :=_dce .(_cfb .EndElement );_gbd &&_daed .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_dce );};func (_ceg ST_CryptProv )String ()string {switch _ceg {case 0:return "";case 1:return "\u0072\u0073\u0061\u0041\u0045\u0053";case 2:return "\u0072s\u0061\u0046\u0075\u006c\u006c";case 3:return "\u0063\u0075\u0073\u0074\u006f\u006d";};return "";};func (_abd *ST_OnOff1 )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_abd =0;case "\u006f\u006e":*_abd =1;case "\u006f\u0066\u0066":*_abd =2;};return nil ;};func (_dg *ST_YAlign )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_dg =0;case "\u0069\u006e\u006c\u0069\u006e\u0065":*_dg =1;case "\u0074\u006f\u0070":*_dg =2;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_dg =3;case "\u0062\u006f\u0074\u0074\u006f\u006d":*_dg =4;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_dg =5;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_dg =6;};return nil ;};func (_edg ST_OnOff1 )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_edg .String (),start );};type ST_CalendarType byte ;type ST_XAlign byte ;func (_ggd ST_XAlign )Validate ()error {return _ggd .ValidateWithPath ("")};func (_cfd ST_AlgClass )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_ac :=_cfb .Attr {};_ac .Name =name ;switch _cfd {case ST_AlgClassUnset :_ac .Value ="";case ST_AlgClassHash :_ac .Value ="\u0068\u0061\u0073\u0068";case ST_AlgClassCustom :_ac .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";};return _ac ,nil ;};const ST_FixedPercentagePattern ="-\u003f\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039]\u003f\u0029\u0029\u0028\u005c\u002e\u005b\u0030\u002d\u0039][\u0030\u002d\u0039]\u003f)\u003f\u0025";var ST_GuidPatternRe =_cd .MustCompile (ST_GuidPattern );func (_bad *ST_CalendarType )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_abe ,_de :=d .Token ();if _de !=nil {return _de ;};if _deb ,_gac :=_abe .(_cfb .EndElement );_gac &&_deb .Name ==start .Name {*_bad =1;return nil ;};if _fgb ,_cb :=_abe .(_cfb .CharData );!_cb {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_abe );}else {switch string (_fgb ){case "":*_bad =0;case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":*_bad =1;case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":*_bad =2;case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":*_bad =3;case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":*_bad =4;case "\u0068\u0069\u006ar\u0069":*_bad =5;case "\u0068\u0065\u0062\u0072\u0065\u0077":*_bad =6;case "\u0074\u0061\u0069\u0077\u0061\u006e":*_bad =7;case "\u006a\u0061\u0070a\u006e":*_bad =8;case "\u0074\u0068\u0061\u0069":*_bad =9;case "\u006b\u006f\u0072e\u0061":*_bad =10;case "\u0073\u0061\u006b\u0061":*_bad =11;case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":*_bad =12;case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":*_bad =13;case "\u006e\u006f\u006e\u0065":*_bad =14;};};_abe ,_de =d .Token ();if _de !=nil {return _de ;};if _cgf ,_aa :=_abe .(_cfb .EndElement );_aa &&_cgf .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_abe );};func ParseUnionST_OnOff (s string )(ST_OnOff ,error ){_bac :=ST_OnOff {};switch s {case "\u0074\u0072\u0075\u0065","\u0031","\u006f\u006e":_ga :=true ;_bac .Bool =&_ga ;default:_fff :=false ;_bac .Bool =&_fff ;};return _bac ,nil ;};func (_ccbb *ST_AlgType )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_eeg ,_fdg :=d .Token ();if _fdg !=nil {return _fdg ;};if _edab ,_df :=_eeg .(_cfb .EndElement );_df &&_edab .Name ==start .Name {*_ccbb =1;return nil ;};if _d ,_edb :=_eeg .(_cfb .CharData );!_edb {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_eeg );}else {switch string (_d ){case "":*_ccbb =0;case "\u0074y\u0070\u0065\u0041\u006e\u0079":*_ccbb =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_ccbb =2;};};_eeg ,_fdg =d .Token ();if _fdg !=nil {return _fdg ;};if _edfe ,_eb :=_eeg .(_cfb .EndElement );_eb &&_edfe .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_eeg );};func (_dda ST_TrueFalseBlank )Validate ()error {return _dda .ValidateWithPath ("")};var ST_PercentagePatternRe =_cd .MustCompile (ST_PercentagePattern );func (_ab *ST_TwipsMeasure )Validate ()error {return _ab .ValidateWithPath ("")};type ST_YAlign byte ;func (_gff ST_OnOff1 )ValidateWithPath (path string )error {switch _gff {case 0,1,2:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_gff ));};return nil ;};func (_cgb ST_VerticalAlignRun )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_cgb .String (),start );};func (_dae ST_TrueFalseBlank )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_dae .String (),start );};func (_ggc ST_XAlign )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_adc :=_cfb .Attr {};_adc .Name =name ;switch _ggc {case ST_XAlignUnset :_adc .Value ="";case ST_XAlignLeft :_adc .Value ="\u006c\u0065\u0066\u0074";case ST_XAlignCenter :_adc .Value ="\u0063\u0065\u006e\u0074\u0065\u0072";case ST_XAlignRight :_adc .Value ="\u0072\u0069\u0067h\u0074";case ST_XAlignInside :_adc .Value ="\u0069\u006e\u0073\u0069\u0064\u0065";case ST_XAlignOutside :_adc .Value ="\u006fu\u0074\u0073\u0069\u0064\u0065";};return _adc ,nil ;};func (_aag *ST_CryptProv )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_aag =0;case "\u0072\u0073\u0061\u0041\u0045\u0053":*_aag =1;case "\u0072s\u0061\u0046\u0075\u006c\u006c":*_aag =2;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_aag =3;};return nil ;};var ST_PositiveUniversalMeasurePatternRe =_cd .MustCompile (ST_PositiveUniversalMeasurePattern );func (_cccb ST_YAlign )String ()string {switch _cccb {case 0:return "";case 1:return "\u0069\u006e\u006c\u0069\u006e\u0065";case 2:return "\u0074\u006f\u0070";case 3:return "\u0063\u0065\u006e\u0074\u0065\u0072";case 4:return "\u0062\u006f\u0074\u0074\u006f\u006d";case 5:return "\u0069\u006e\u0073\u0069\u0064\u0065";case 6:return "\u006fu\u0074\u0073\u0069\u0064\u0065";};return "";};func (_dfcd ST_AlgType )Validate ()error {return _dfcd .ValidateWithPath ("")};func (_ceb ST_TrueFalse )ValidateWithPath (path string )error {switch _ceb {case 0,1,2,3,4:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_ceb ));};return nil ;};func (_fgeb ST_TrueFalseBlank )String ()string {switch _fgeb {case 0:return "";case 1:return "\u0074";case 2:return "\u0066";case 3:return "\u0074\u0072\u0075\u0065";case 4:return "\u0066\u0061\u006cs\u0065";case 6:return "\u0054\u0072\u0075\u0065";case 7:return "\u0046\u0061\u006cs\u0065";};return "";};func (_gef *ST_TrueFalse )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_gef =0;case "\u0074":*_gef =1;case "\u0066":*_gef =2;case "\u0074\u0072\u0075\u0065":*_gef =3;case "\u0066\u0061\u006cs\u0065":*_gef =4;};return nil ;};func (_bg ST_CryptProv )Validate ()error {return _bg .ValidateWithPath ("")};const (ST_CryptProvUnset ST_CryptProv =0;ST_CryptProvRsaAES ST_CryptProv =1;ST_CryptProvRsaFull ST_CryptProv =2;ST_CryptProvCustom ST_CryptProv =3;);

// ST_TwipsMeasure is a union type
type ST_TwipsMeasure struct{ST_UnsignedDecimalNumber *uint64 ;ST_PositiveUniversalMeasure *string ;};func (_ag *ST_AlgType )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_ag =0;case "\u0074y\u0070\u0065\u0041\u006e\u0079":*_ag =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_ag =2;};return nil ;};func (_fgc ST_ConformanceClass )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_fgc .String (),start );};func (_faa *ST_OnOff )ValidateWithPath (path string )error {_be :=[]string {};if _faa .Bool !=nil {_be =append (_be ,"\u0042\u006f\u006f\u006c");};if _faa .ST_OnOff1 !=ST_OnOff1Unset {_be =append (_be ,"\u0053T\u005f\u004f\u006e\u004f\u0066\u00661");};if len (_be )> 1{return _dfg .Errorf ("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076",path ,_be );};return nil ;};func (_abca ST_TrueFalse )Validate ()error {return _abca .ValidateWithPath ("")};func (_dfc ST_AlgType )String ()string {switch _dfc {case 0:return "";case 1:return "\u0074y\u0070\u0065\u0041\u006e\u0079";case 2:return "\u0063\u0075\u0073\u0074\u006f\u006d";};return "";};func (_b *ST_XAlign )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_dbd ,_ffa :=d .Token ();if _ffa !=nil {return _ffa ;};if _bbc ,_bfa :=_dbd .(_cfb .EndElement );_bfa &&_bbc .Name ==start .Name {*_b =1;return nil ;};if _gdf ,_eaf :=_dbd .(_cfb .CharData );!_eaf {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_dbd );}else {switch string (_gdf ){case "":*_b =0;case "\u006c\u0065\u0066\u0074":*_b =1;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_b =2;case "\u0072\u0069\u0067h\u0074":*_b =3;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_b =4;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_b =5;};};_dbd ,_ffa =d .Token ();if _ffa !=nil {return _ffa ;};if _edf ,_fad :=_dbd .(_cfb .EndElement );_fad &&_edf .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_dbd );};func (_egf *ST_OnOff )Validate ()error {return _egf .ValidateWithPath ("")};func (_a *ST_TrueFalseBlank )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_a =0;case "\u0074":*_a =1;case "\u0066":*_a =2;case "\u0074\u0072\u0075\u0065":*_a =3;case "\u0066\u0061\u006cs\u0065":*_a =4;case "\u0054\u0072\u0075\u0065":*_a =6;case "\u0046\u0061\u006cs\u0065":*_a =7;};return nil ;};func (_cbf ST_AlgClass )String ()string {switch _cbf {case 0:return "";case 1:return "\u0068\u0061\u0073\u0068";case 2:return "\u0063\u0075\u0073\u0074\u006f\u006d";};return "";};const (ST_YAlignUnset ST_YAlign =0;ST_YAlignInline ST_YAlign =1;ST_YAlignTop ST_YAlign =2;ST_YAlignCenter ST_YAlign =3;ST_YAlignBottom ST_YAlign =4;ST_YAlignInside ST_YAlign =5;ST_YAlignOutside ST_YAlign =6;);func (_gb ST_ConformanceClass )Validate ()error {return _gb .ValidateWithPath ("")};const (ST_TrueFalseBlankUnset ST_TrueFalseBlank =0;ST_TrueFalseBlankT ST_TrueFalseBlank =1;ST_TrueFalseBlankF ST_TrueFalseBlank =2;ST_TrueFalseBlankTrue ST_TrueFalseBlank =3;ST_TrueFalseBlankFalse ST_TrueFalseBlank =4;ST_TrueFalseBlankTrue_ ST_TrueFalseBlank =6;ST_TrueFalseBlankFalse_ ST_TrueFalseBlank =7;);func (_bgg ST_YAlign )Validate ()error {return _bgg .ValidateWithPath ("")};func (_cef *ST_AlgClass )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_dgf ,_ddb :=d .Token ();if _ddb !=nil {return _ddb ;};if _cca ,_bfac :=_dgf .(_cfb .EndElement );_bfac &&_cca .Name ==start .Name {*_cef =1;return nil ;};if _ccc ,_fdd :=_dgf .(_cfb .CharData );!_fdd {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_dgf );}else {switch string (_ccc ){case "":*_cef =0;case "\u0068\u0061\u0073\u0068":*_cef =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_cef =2;};};_dgf ,_ddb =d .Token ();if _ddb !=nil {return _ddb ;};if _ge ,_dad :=_dgf .(_cfb .EndElement );_dad &&_ge .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_dgf );};func (_ea ST_TrueFalse )String ()string {switch _ea {case 0:return "";case 1:return "\u0074";case 2:return "\u0066";case 3:return "\u0074\u0072\u0075\u0065";case 4:return "\u0066\u0061\u006cs\u0065";};return "";};type ST_VerticalAlignRun byte ;func (_ced ST_ConformanceClass )ValidateWithPath (path string )error {switch _ced {case 0,1,2:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_ced ));};return nil ;};const (ST_VerticalAlignRunUnset ST_VerticalAlignRun =0;ST_VerticalAlignRunBaseline ST_VerticalAlignRun =1;ST_VerticalAlignRunSuperscript ST_VerticalAlignRun =2;ST_VerticalAlignRunSubscript ST_VerticalAlignRun =3;);func (_gbf *ST_ConformanceClass )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_gbf =0;case "\u0073\u0074\u0072\u0069\u0063\u0074":*_gbf =1;case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":*_gbf =2;};return nil ;};var ST_PositivePercentagePatternRe =_cd .MustCompile (ST_PositivePercentagePattern );func (_fcg ST_ConformanceClass )String ()string {switch _fcg {case 0:return "";case 1:return "\u0073\u0074\u0072\u0069\u0063\u0074";case 2:return "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c";};return "";};func (_gcd ST_CryptProv )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_gcd .String (),start );};func (_cfe ST_AlgClass )ValidateWithPath (path string )error {switch _cfe {case 0,1,2:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_cfe ));};return nil ;};func (_g ST_CalendarType )ValidateWithPath (path string )error {switch _g {case 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_g ));};return nil ;};const ST_UniversalMeasurePattern ="\u002d\u003f\u005b\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e\u005b\u0030\u002d\u0039\u005d\u002b\u0029\u003f\u0028\u006d\u006d\u007c\u0063m\u007c\u0069\u006e\u007c\u0070t\u007c\u0070c\u007c\u0070\u0069\u0029";func (_ba *ST_CalendarType )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_ba =0;case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":*_ba =1;case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":*_ba =2;case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":*_ba =3;case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":*_ba =4;case "\u0068\u0069\u006ar\u0069":*_ba =5;case "\u0068\u0065\u0062\u0072\u0065\u0077":*_ba =6;case "\u0074\u0061\u0069\u0077\u0061\u006e":*_ba =7;case "\u006a\u0061\u0070a\u006e":*_ba =8;case "\u0074\u0068\u0061\u0069":*_ba =9;case "\u006b\u006f\u0072e\u0061":*_ba =10;case "\u0073\u0061\u006b\u0061":*_ba =11;case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":*_ba =12;case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":*_ba =13;case "\u006e\u006f\u006e\u0065":*_ba =14;};return nil ;};func (_bb ST_TwipsMeasure )String ()string {if _bb .ST_UnsignedDecimalNumber !=nil {return _dfg .Sprintf ("\u0025\u0076",*_bb .ST_UnsignedDecimalNumber );};if _bb .ST_PositiveUniversalMeasure !=nil {return _dfg .Sprintf ("\u0025\u0076",*_bb .ST_PositiveUniversalMeasure );};return "";};func (_cea ST_TrueFalse )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_ff :=_cfb .Attr {};_ff .Name =name ;switch _cea {case ST_TrueFalseUnset :_ff .Value ="";case ST_TrueFalseT :_ff .Value ="\u0074";case ST_TrueFalseF :_ff .Value ="\u0066";case ST_TrueFalseTrue :_ff .Value ="\u0074\u0072\u0075\u0065";case ST_TrueFalseFalse :_ff .Value ="\u0066\u0061\u006cs\u0065";};return _ff ,nil ;};func (_c ST_VerticalAlignRun )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_gdb :=_cfb .Attr {};_gdb .Name =name ;switch _c {case ST_VerticalAlignRunUnset :_gdb .Value ="";case ST_VerticalAlignRunBaseline :_gdb .Value ="\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065";case ST_VerticalAlignRunSuperscript :_gdb .Value ="s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074";case ST_VerticalAlignRunSubscript :_gdb .Value ="\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t";};return _gdb ,nil ;};func (_ad ST_VerticalAlignRun )String ()string {switch _ad {case 0:return "";case 1:return "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065";case 2:return "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074";case 3:return "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t";};return "";};func (_fcd ST_TrueFalseBlank )ValidateWithPath (path string )error {switch _fcd {case 0,1,2,3,4,6,7:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_fcd ));};return nil ;};func (_fcc *ST_TwipsMeasure )ValidateWithPath (path string )error {_ae :=[]string {};if _fcc .ST_UnsignedDecimalNumber !=nil {_ae =append (_ae ,"\u0053T\u005f\u0055\u006e\u0073\u0069\u0067\u006e\u0065\u0064\u0044\u0065c\u0069\u006d\u0061\u006c\u004e\u0075\u006d\u0062\u0065\u0072");};if _fcc .ST_PositiveUniversalMeasure !=nil {_ae =append (_ae ,"S\u0054\u005f\u0050\u006f\u0073\u0069t\u0069\u0076\u0065\u0055\u006e\u0069\u0076\u0065\u0072s\u0061\u006c\u004de\u0061s\u0075\u0072\u0065");};if len (_ae )> 1{return _dfg .Errorf ("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076",path ,_ae );};return nil ;};func (_acc *ST_YAlign )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_cba ,_cfdg :=d .Token ();if _cfdg !=nil {return _cfdg ;};if _ed ,_ffe :=_cba .(_cfb .EndElement );_ffe &&_ed .Name ==start .Name {*_acc =1;return nil ;};if _gd ,_cdf :=_cba .(_cfb .CharData );!_cdf {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_cba );}else {switch string (_gd ){case "":*_acc =0;case "\u0069\u006e\u006c\u0069\u006e\u0065":*_acc =1;case "\u0074\u006f\u0070":*_acc =2;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_acc =3;case "\u0062\u006f\u0074\u0074\u006f\u006d":*_acc =4;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_acc =5;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_acc =6;};};_cba ,_cfdg =d .Token ();if _cfdg !=nil {return _cfdg ;};if _bdg ,_baf :=_cba .(_cfb .EndElement );_baf &&_bdg .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_cba );};func (_ecg ST_TwipsMeasure )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {e .EncodeToken (start );if _ecg .ST_UnsignedDecimalNumber !=nil {e .EncodeToken (_cfb .CharData (_dfg .Sprintf ("\u0025\u0064",*_ecg .ST_UnsignedDecimalNumber )));};if _ecg .ST_PositiveUniversalMeasure !=nil {e .EncodeToken (_cfb .CharData (*_ecg .ST_PositiveUniversalMeasure ));};return e .EncodeToken (_cfb .EndElement {Name :start .Name });};func (_egd ST_CalendarType )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_fb :=_cfb .Attr {};_fb .Name =name ;switch _egd {case ST_CalendarTypeUnset :_fb .Value ="";case ST_CalendarTypeGregorian :_fb .Value ="\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n";case ST_CalendarTypeGregorianUs :_fb .Value ="g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073";case ST_CalendarTypeGregorianMeFrench :_fb .Value ="\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068";case ST_CalendarTypeGregorianArabic :_fb .Value ="\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063";case ST_CalendarTypeHijri :_fb .Value ="\u0068\u0069\u006ar\u0069";case ST_CalendarTypeHebrew :_fb .Value ="\u0068\u0065\u0062\u0072\u0065\u0077";case ST_CalendarTypeTaiwan :_fb .Value ="\u0074\u0061\u0069\u0077\u0061\u006e";case ST_CalendarTypeJapan :_fb .Value ="\u006a\u0061\u0070a\u006e";case ST_CalendarTypeThai :_fb .Value ="\u0074\u0068\u0061\u0069";case ST_CalendarTypeKorea :_fb .Value ="\u006b\u006f\u0072e\u0061";case ST_CalendarTypeSaka :_fb .Value ="\u0073\u0061\u006b\u0061";case ST_CalendarTypeGregorianXlitEnglish :_fb .Value ="g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068";case ST_CalendarTypeGregorianXlitFrench :_fb .Value ="\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068";case ST_CalendarTypeNone :_fb .Value ="\u006e\u006f\u006e\u0065";};return _fb ,nil ;};func (_add *ST_CryptProv )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_eg ,_bfc :=d .Token ();if _bfc !=nil {return _bfc ;};if _bd ,_fge :=_eg .(_cfb .EndElement );_fge &&_bd .Name ==start .Name {*_add =1;return nil ;};if _aac ,_fbf :=_eg .(_cfb .CharData );!_fbf {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_eg );}else {switch string (_aac ){case "":*_add =0;case "\u0072\u0073\u0061\u0041\u0045\u0053":*_add =1;case "\u0072s\u0061\u0046\u0075\u006c\u006c":*_add =2;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_add =3;};};_eg ,_bfc =d .Token ();if _bfc !=nil {return _bfc ;};if _fcf ,_fc :=_eg .(_cfb .EndElement );_fc &&_fcf .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_eg );};func (_gdfe ST_OnOff1 )Validate ()error {return _gdfe .ValidateWithPath ("")};func (_dd ST_CryptProv )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_db :=_cfb .Attr {};_db .Name =name ;switch _dd {case ST_CryptProvUnset :_db .Value ="";case ST_CryptProvRsaAES :_db .Value ="\u0072\u0073\u0061\u0041\u0045\u0053";case ST_CryptProvRsaFull :_db .Value ="\u0072s\u0061\u0046\u0075\u006c\u006c";case ST_CryptProvCustom :_db .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";};return _db ,nil ;};func (_cc *ST_TrueFalse )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_eccg ,_bdd :=d .Token ();if _bdd !=nil {return _bdd ;};if _bc ,_abc :=_eccg .(_cfb .EndElement );_abc &&_bc .Name ==start .Name {*_cc =1;return nil ;};if _gfa ,_gg :=_eccg .(_cfb .CharData );!_gg {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_eccg );}else {switch string (_gfa ){case "":*_cc =0;case "\u0074":*_cc =1;case "\u0066":*_cc =2;case "\u0074\u0072\u0075\u0065":*_cc =3;case "\u0066\u0061\u006cs\u0065":*_cc =4;};};_eccg ,_bdd =d .Token ();if _bdd !=nil {return _bdd ;};if _ec ,_ccbe :=_eccg .(_cfb .EndElement );_ccbe &&_ec .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_eccg );};func (_ebg *ST_VerticalAlignRun )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_ebg =0;case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":*_ebg =1;case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":*_ebg =2;case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":*_ebg =3;};return nil ;};func (_bgd *ST_ConformanceClass )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_gcf ,_gcg :=d .Token ();if _gcg !=nil {return _gcg ;};if _dfe ,_dga :=_gcf .(_cfb .EndElement );_dga &&_dfe .Name ==start .Name {*_bgd =1;return nil ;};if _ebd ,_ee :=_gcf .(_cfb .CharData );!_ee {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gcf );}else {switch string (_ebd ){case "":*_bgd =0;case "\u0073\u0074\u0072\u0069\u0063\u0074":*_bgd =1;case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":*_bgd =2;};};_gcf ,_gcg =d .Token ();if _gcg !=nil {return _gcg ;};if _ca ,_cae :=_gcf .(_cfb .EndElement );_cae &&_ca .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gcf );};type ST_OnOff1 byte ;func (_fe *ST_TrueFalseBlank )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_gdg ,_aad :=d .Token ();if _aad !=nil {return _aad ;};if _fg ,_cf :=_gdg .(_cfb .EndElement );_cf &&_fg .Name ==start .Name {*_fe =1;return nil ;};if _abf ,_cfg :=_gdg .(_cfb .CharData );!_cfg {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gdg );}else {switch string (_abf ){case "":*_fe =0;case "\u0074":*_fe =1;case "\u0066":*_fe =2;case "\u0074\u0072\u0075\u0065":*_fe =3;case "\u0066\u0061\u006cs\u0065":*_fe =4;case "\u0054\u0072\u0075\u0065":*_fe =6;case "\u0046\u0061\u006cs\u0065":*_fe =7;};};_gdg ,_aad =d .Token ();if _aad !=nil {return _aad ;};if _gbc ,_cag :=_gdg .(_cfb .EndElement );_cag &&_gbc .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gdg );};func (_eac ST_CalendarType )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_eac .String (),start );};type ST_ConformanceClass byte ;func (_cg ST_CryptProv )ValidateWithPath (path string )error {switch _cg {case 0,1,2,3:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_cg ));};return nil ;};func (_dcea ST_OnOff )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {e .EncodeToken (start );if _dcea .Bool !=nil {e .EncodeToken (_cfb .CharData (_dfg .Sprintf ("\u0025\u0064",_cdg (*_dcea .Bool ))));};if _dcea .ST_OnOff1 !=ST_OnOff1Unset {e .EncodeToken (_cfb .CharData (_dcea .ST_OnOff1 .String ()));};return e .EncodeToken (_cfb .EndElement {Name :start .Name });};func (_ce ST_AlgType )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_cagd :=_cfb .Attr {};_cagd .Name =name ;switch _ce {case ST_AlgTypeUnset :_cagd .Value ="";case ST_AlgTypeTypeAny :_cagd .Value ="\u0074y\u0070\u0065\u0041\u006e\u0079";case ST_AlgTypeCustom :_cagd .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";};return _cagd ,nil ;};func (_aaf ST_TrueFalse )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_aaf .String (),start );};func (_cbd ST_VerticalAlignRun )ValidateWithPath (path string )error {switch _cbd {case 0,1,2,3:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_cbd ));};return nil ;};const ST_GuidPattern ="\u005c\u007b\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b\u0038\u007d\u002d\u005b\u0030\u002d9\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030-\u0039\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b4\u007d\u002d\u005b\u0030\u002d\u0039A\u002d\u0046]\u007b\u00312\u007d\\\u007d";const ST_PercentagePattern ="-\u003f[\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e[\u0030\u002d\u0039\u005d+)\u003f\u0025";func (_ecc ST_OnOff1 )String ()string {switch _ecc {case 0:return "";case 1:return "\u006f\u006e";case 2:return "\u006f\u0066\u0066";};return "";};type ST_TrueFalse byte ;const (ST_CalendarTypeUnset ST_CalendarType =0;ST_CalendarTypeGregorian ST_CalendarType =1;ST_CalendarTypeGregorianUs ST_CalendarType =2;ST_CalendarTypeGregorianMeFrench ST_CalendarType =3;ST_CalendarTypeGregorianArabic ST_CalendarType =4;ST_CalendarTypeHijri ST_CalendarType =5;ST_CalendarTypeHebrew ST_CalendarType =6;ST_CalendarTypeTaiwan ST_CalendarType =7;ST_CalendarTypeJapan ST_CalendarType =8;ST_CalendarTypeThai ST_CalendarType =9;ST_CalendarTypeKorea ST_CalendarType =10;ST_CalendarTypeSaka ST_CalendarType =11;ST_CalendarTypeGregorianXlitEnglish ST_CalendarType =12;ST_CalendarTypeGregorianXlitFrench ST_CalendarType =13;ST_CalendarTypeNone ST_CalendarType =14;);func (_dcd ST_YAlign )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_dcd .String (),start );};const ST_PositiveUniversalMeasurePattern ="\u005b\u0030-9\u005d\u002b\u0028\\\u002e\u005b\u0030\u002d9]+\u0029?(\u006d\u006d\u007c\u0063\u006d\u007c\u0069n|\u0070\u0074\u007c\u0070\u0063\u007c\u0070i\u0029";func (_dbe ST_VerticalAlignRun )Validate ()error {return _dbe .ValidateWithPath ("")};const (ST_OnOff1Unset ST_OnOff1 =0;ST_OnOff1On ST_OnOff1 =1;ST_OnOff1Off ST_OnOff1 =2;);func (_afg *ST_XAlign )UnmarshalXMLAttr (attr _cfb .Attr )error {switch attr .Value {case "":*_afg =0;case "\u006c\u0065\u0066\u0074":*_afg =1;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_afg =2;case "\u0072\u0069\u0067h\u0074":*_afg =3;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_afg =4;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_afg =5;};return nil ;};func (_eef ST_AlgClass )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_eef .String (),start );};const (ST_AlgTypeUnset ST_AlgType =0;ST_AlgTypeTypeAny ST_AlgType =1;ST_AlgTypeCustom ST_AlgType =2;);func (_dc ST_OnOff1 )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_gdd :=_cfb .Attr {};_gdd .Name =name ;switch _dc {case ST_OnOff1Unset :_gdd .Value ="";case ST_OnOff1On :_gdd .Value ="\u006f\u006e";case ST_OnOff1Off :_gdd .Value ="\u006f\u0066\u0066";};return _gdd ,nil ;};

// ST_OnOff is a union type
type ST_OnOff struct{Bool *bool ;ST_OnOff1 ST_OnOff1 ;};var ST_FixedPercentagePatternRe =_cd .MustCompile (ST_FixedPercentagePattern );func (_gdga ST_XAlign )ValidateWithPath (path string )error {switch _gdga {case 0,1,2,3,4,5:default:return _dfg .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_gdga ));};return nil ;};func (_af ST_OnOff )String ()string {if _af .Bool !=nil {return _dfg .Sprintf ("\u0025\u0076",*_af .Bool );};if _af .ST_OnOff1 !=ST_OnOff1Unset {return _af .ST_OnOff1 .String ();};return "";};func (_da ST_CalendarType )String ()string {switch _da {case 0:return "";case 1:return "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n";case 2:return "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073";case 3:return "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068";case 4:return "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063";case 5:return "\u0068\u0069\u006ar\u0069";case 6:return "\u0068\u0065\u0062\u0072\u0065\u0077";case 7:return "\u0074\u0061\u0069\u0077\u0061\u006e";case 8:return "\u006a\u0061\u0070a\u006e";case 9:return "\u0074\u0068\u0061\u0069";case 10:return "\u006b\u006f\u0072e\u0061";case 11:return "\u0073\u0061\u006b\u0061";case 12:return "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068";case 13:return "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068";case 14:return "\u006e\u006f\u006e\u0065";};return "";};func (_feb *ST_OnOff1 )UnmarshalXML (d *_cfb .Decoder ,start _cfb .StartElement )error {_ccb ,_dcb :=d .Token ();if _dcb !=nil {return _dcb ;};if _ega ,_dbgg :=_ccb .(_cfb .EndElement );_dbgg &&_ega .Name ==start .Name {*_feb =1;return nil ;};if _eda ,_cfbg :=_ccb .(_cfb .CharData );!_cfbg {return _dfg .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_ccb );}else {switch string (_eda ){case "":*_feb =0;case "\u006f\u006e":*_feb =1;case "\u006f\u0066\u0066":*_feb =2;};};_ccb ,_dcb =d .Token ();if _dcb !=nil {return _dcb ;};if _agg ,_fca :=_ccb .(_cfb .EndElement );_fca &&_agg .Name ==start .Name {return nil ;};return _dfg .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_ccb );};type ST_CryptProv byte ;func (_fd ST_ConformanceClass )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_eaa :=_cfb .Attr {};_eaa .Name =name ;switch _fd {case ST_ConformanceClassUnset :_eaa .Value ="";case ST_ConformanceClassStrict :_eaa .Value ="\u0073\u0074\u0072\u0069\u0063\u0074";case ST_ConformanceClassTransitional :_eaa .Value ="\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c";};return _eaa ,nil ;};func (_gc ST_AlgType )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_gc .String (),start );};const (ST_AlgClassUnset ST_AlgClass =0;ST_AlgClassHash ST_AlgClass =1;ST_AlgClassCustom ST_AlgClass =2;);func (_dbb ST_TrueFalseBlank )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_gf :=_cfb .Attr {};_gf .Name =name ;switch _dbb {case ST_TrueFalseBlankUnset :_gf .Value ="";case ST_TrueFalseBlankT :_gf .Value ="\u0074";case ST_TrueFalseBlankF :_gf .Value ="\u0066";case ST_TrueFalseBlankTrue :_gf .Value ="\u0074\u0072\u0075\u0065";case ST_TrueFalseBlankFalse :_gf .Value ="\u0066\u0061\u006cs\u0065";case ST_TrueFalseBlankTrue_ :_gf .Value ="\u0054\u0072\u0075\u0065";case ST_TrueFalseBlankFalse_ :_gf .Value ="\u0046\u0061\u006cs\u0065";};return _gf ,nil ;};func (_afa ST_XAlign )MarshalXML (e *_cfb .Encoder ,start _cfb .StartElement )error {return e .EncodeElement (_afa .String (),start );};func (_fa ST_YAlign )MarshalXMLAttr (name _cfb .Name )(_cfb .Attr ,error ){_dcg :=_cfb .Attr {};_dcg .Name =name ;switch _fa {case ST_YAlignUnset :_dcg .Value ="";case ST_YAlignInline :_dcg .Value ="\u0069\u006e\u006c\u0069\u006e\u0065";case ST_YAlignTop :_dcg .Value ="\u0074\u006f\u0070";case ST_YAlignCenter :_dcg .Value ="\u0063\u0065\u006e\u0074\u0065\u0072";case ST_YAlignBottom :_dcg .Value ="\u0062\u006f\u0074\u0074\u006f\u006d";case ST_YAlignInside :_dcg .Value ="\u0069\u006e\u0073\u0069\u0064\u0065";case ST_YAlignOutside :_dcg .Value ="\u006fu\u0074\u0073\u0069\u0064\u0065";};return _dcg ,nil ;};type ST_AlgClass byte ;func _cdg (_adf bool )uint8 {if _adf {return 1;};return 0;};