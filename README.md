# laba2_soa

Run: go run main.go

Results:
Go serialization
Time serialization:  678.292µs
Time deserialization:  60.541µs
Length string:  188
Output:  R��message��IDStrData
                                FloatData       DataArray��DataMap���[]string��
                                                                               ��map[string]int��
                                                                                                 2��test�@5^�?132testrsdftest1test2

Json serialization
Time serialization:  72.208µs
Time deserialization:  37.375µs
Length string:  109
Output:  {"id":1,"strData":"test","floatData":1.023,"dataArray":["132","test","rsdf"],"dataMap":{"test1":1,"test2":2}}

Yaml serialization
Time serialization:  61.625µs
Time deserialization:  73.333µs
Length string:  101
Output:  id: 1
strdata: test
floatdata: 1.023
dataarray:
- "132"
- test
- rsdf
datamap:
  test1: 1
  test2: 2


MsgPack serialization
Time serialization:  31.291µs
Time deserialization:  24.417µs
Length string:  81
Output:  ��ID�StrData�test�FloatData�?��DataArray��132�test�rsdf�DataMap��test1�test2

Avro serialization
Time serialization:  148.083µs
Time deserialization:  27.625µs
Length string:  44
Output: test��?13tesrsdf
test1
test2

Protobuf serialization
Time serialization:  148.542µs
Time deserialization:  12.084µs
Length string:  47
Output: test"132"test"rsdf*     
test2*  
test1

XML serialization
Time serialization:  33.875µs
Time deserialization:  37.208µs
Length string:  166
Output:  <messageXml><ID>1</ID><StrData>test</StrData><FloatData>1.023</FloatData><DataArray>132</DataArray><DataArray>test</DataArray><DataArray>rsdf</DataArray></messageXml>

