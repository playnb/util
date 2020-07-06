#Util工具箱

###log
初始化默认的日志
<pre>
	log.InitPanic("../tmp")
	log.Init(log.DefaultLogger("../tmp", "run"))
	defer log.Flush()
</pre>

----

###BuffData
使用默认的缓冲池
<pre>
util.DefaultPool()
</pre>
获取缓冲数据对象(长度800)
<pre>
data:=util.DefaultPool().Get(800)
</pre>
归还缓冲数据对象
<pre>
data.Release()
</pre>