var libFuncName=null;if(typeof jQuery=="undefined"&&typeof Zepto=="undefined"&&typeof $=="function"){libFuncName=$}else{if(typeof jQuery=="function"){libFuncName=jQuery}else{if(typeof Zepto!="function"){throw new TypeError}libFuncName=Zepto}}(function(c,a,d,b){a.matchMedia=a.matchMedia||function(k,g){var m,j=k.documentElement,f=j.firstElementChild||j.firstChild,h=k.createElement("body"),l=k.createElement("div");return l.id="mq-test-1",l.style.cssText="position:absolute;top:-100em",h.style.background="none",h.appendChild(l),function(i){return l.innerHTML='&shy;<style media="'+i+'"> #mq-test-1 { width: 42px; }</style>',j.insertBefore(h,f),m=l.offsetWidth===42,j.removeChild(h),{matches:m,media:i}}}(d),Array.prototype.filter||(Array.prototype.filter=function(k){if(this==null){throw new TypeError}var g=Object(this),m=g.length>>>0;if(typeof k!="function"){return}var j=[],f=arguments[1];for(var h=0;h<m;h++){if(h in g){var l=g[h];k&&k.call(f,l,h,g)&&j.push(l)}}return j}),Function.prototype.bind||(Function.prototype.bind=function(j){if(typeof this!="function"){throw new TypeError("Function.prototype.bind - what is trying to be bound is not callable")}var g=Array.prototype.slice.call(arguments,1),k=this,h=function(){},f=function(){return k.apply(this instanceof h&&j?this:j,g.concat(Array.prototype.slice.call(arguments)))};return h.prototype=this.prototype,f.prototype=new h,f}),Array.prototype.indexOf||(Array.prototype.indexOf=function(j){if(this==null){throw new TypeError}var g=Object(this),k=g.length>>>0;if(k===0){return -1}var h=0;arguments.length>1&&(h=Number(arguments[1]),h!=h?h=0:h!=0&&h!=Infinity&&h!=-Infinity&&(h=(h>0||-1)*Math.floor(Math.abs(h))));if(h>=k){return -1}var f=h>=0?h:Math.max(k-Math.abs(h),0);for(;f<k;f++){if(f in g&&g[f]===j){return f}}return -1}),c.fn.stop=c.fn.stop||function(){return this},a.Foundation={name:"Foundation",version:"4.2.0",cache:{},init:function(w,h,e,k,x,g){var v,q=[w,e,k,x],m=[],g=g||!1;g&&(this.nc=g),this.rtl=/rtl/i.test(c("html").attr("dir")),this.scope=w||this.scope;if(h&&typeof h=="string"&&!/reflow/i.test(h)){if(/off/i.test(h)){return this.off()}v=h.split(" ");if(v.length>0){for(var j=v.length-1;j>=0;j--){m.push(this.init_lib(v[j],q))}}}else{/reflow/i.test(h)&&(q[1]="reflow");for(var p in this.libs){m.push(this.init_lib(p,q))}}return typeof h=="function"&&q.unshift(h),this.response_obj(m,q)},response_obj:function(h,f){for(var i=0,g=f.length;i<g;i++){if(typeof f[i]=="function"){return f[i]({errors:h.filter(function(j){if(typeof j=="string"){return j}})})}}return h},init_lib:function(g,f){return this.trap(function(){return this.libs.hasOwnProperty(g)?(this.patch(this.libs[g]),this.libs[g].init.apply(this.libs[g],f)):function(){}}.bind(this),g)},trap:function(g,f){if(!this.nc){try{return g()}catch(h){return this.error({name:f,message:"could not be initialized",more:h.name+" "+h.message})}}return g()},patch:function(f){this.fix_outer(f),f.scope=this.scope,f.rtl=this.rtl},inherit:function(h,f){var i=f.split(" ");for(var g=i.length-1;g>=0;g--){this.lib_methods.hasOwnProperty(i[g])&&(this.libs[h.name][i[g]]=this.lib_methods[i[g]])}},random_str:function(h){var f="0123456789ABCDEFGHIJKLMNOPQRSTUVWXTZabcdefghiklmnopqrstuvwxyz".split("");h||(h=Math.floor(Math.random()*f.length));var i="";for(var g=0;g<h;g++){i+=f[Math.floor(Math.random()*f.length)]}return i},libs:{},lib_methods:{set_data:function(g,f){var h=[this.name,+(new Date),Foundation.random_str(5)].join("-");return Foundation.cache[h]=f,g.attr("data-"+this.name+"-id",h),f},get_data:function(f){return Foundation.cache[f.attr("data-"+this.name+"-id")]},remove_data:function(e){e?(delete Foundation.cache[e.attr("data-"+this.name+"-id")],e.attr("data-"+this.name+"-id","")):c("[data-"+this.name+"-id]").each(function(){delete Foundation.cache[c(this).attr("data-"+this.name+"-id")],c(this).attr("data-"+this.name+"-id","")})},throttle:function(g,f){var h=null;return function(){var j=this,e=arguments;clearTimeout(h),h=setTimeout(function(){g.apply(j,e)},f)}},data_options:function(h){function f(i){return !isNaN(i-0)&&i!==null&&i!==""&&i!==!1&&i!==!0}function e(i){return typeof i=="string"?c.trim(i):i}var m={},k,g,j=(h.attr("data-options")||":").split(";"),l=j.length;for(k=l-1;k>=0;k--){g=j[k].split(":"),/true/i.test(g[1])&&(g[1]=!0),/false/i.test(g[1])&&(g[1]=!1),f(g[1])&&(g[1]=parseInt(g[1],10)),g.length===2&&g[0].length>0&&(m[e(g[0])]=e(g[1]))}return m},delay:function(g,f){return setTimeout(g,f)},scrollTo:function(j,g,e){if(e<0){return}var f=g-c(a).scrollTop(),h=f/e*10;this.scrollToTimerCache=setTimeout(function(){isNaN(parseInt(h,10))||(a.scrollTo(0,c(a).scrollTop()+h),this.scrollTo(j,g,e-10))}.bind(this),10)},scrollLeft:function(f){if(!f.length){return}return"scrollLeft" in f[0]?f[0].scrollLeft:f[0].pageXOffset},empty:function(g){if(g.length&&g.length>0){return !1}if(g.length&&g.length===0){return !0}for(var f in g){if(hasOwnProperty.call(g,f)){return !1}}return !0}},fix_outer:function(f){f.outerHeight=function(h,g){return typeof Zepto=="function"?h.height():typeof g!="undefined"?h.outerHeight(g):h.outerHeight()},f.outerWidth=function(g){return typeof Zepto=="function"?g.width():typeof bool!="undefined"?g.outerWidth(bool):g.outerWidth()}},error:function(f){return f.name+" "+f.message+"; "+f.more},off:function(){return c(this.scope).off(".fndtn"),c(a).off(".fndtn"),!0},zj:function(){return typeof Zepto!="undefined"?Zepto:jQuery}()},c.fn.foundation=function(){var f=Array.prototype.slice.call(arguments,0);return this.each(function(){return Foundation.init.apply(Foundation,[this].concat(f)),this})}})(libFuncName,this,this.document),function(c,a,d,b){Foundation.libs.dropdown={name:"dropdown",version:"4.2.0",settings:{activeClass:"open",is_hover:!1,opened:function(){},closed:function(){}},init:function(e,g,f){return this.scope=e||this.scope,Foundation.inherit(this,"throttle scrollLeft data_options"),typeof g=="object"&&c.extend(!0,this.settings,g),typeof g!="string"?(this.settings.init||this.events(),this.settings.init):this[g].call(this,f)},events:function(){var e=this;c(this.scope).on("click.fndtn.dropdown","[data-dropdown]",function(f){var g=c.extend({},e.settings,e.data_options(c(this)));f.preventDefault(),g.is_hover||e.toggle(c(this))}).on("mouseenter","[data-dropdown]",function(f){var g=c.extend({},e.settings,e.data_options(c(this)));g.is_hover&&e.toggle(c(this))}).on("mouseleave","[data-dropdown-content]",function(g){var h=c('[data-dropdown="'+c(this).attr("id")+'"]'),f=c.extend({},e.settings,e.data_options(h));f.is_hover&&e.close.call(e,c(this))}).on("opened.fndtn.dropdown","[data-dropdown-content]",this.settings.opened).on("closed.fndtn.dropdown","[data-dropdown-content]",this.settings.closed),c("body").on("click.fndtn.dropdown",function(f){var g=c(f.target).closest("[data-dropdown-content]");if(c(f.target).data("dropdown")){return}if(g.length>0&&(c(f.target).is("[data-dropdown-content]")||c.contains(g.first()[0],f.target))){f.stopPropagation();return}e.close.call(e,c("[data-dropdown-content]"))}),c(a).on("resize.fndtn.dropdown",e.throttle(function(){e.resize.call(e)},50)).trigger("resize"),this.settings.init=!0},close:function(e){var f=this;e.each(function(){c(this).hasClass(f.settings.activeClass)&&(c(this).css(Foundation.rtl?"right":"left","-99999px").removeClass(f.settings.activeClass),c(this).trigger("closed"))})},open:function(g,f){this.css(g.addClass(this.settings.activeClass),f),g.trigger("opened")},toggle:function(e){var f=c("#"+e.data("dropdown"));this.close.call(this,c("[data-dropdown-content]").not(f)),f.hasClass(this.settings.activeClass)?this.close.call(this,f):(this.close.call(this,c("[data-dropdown-content]")),this.open.call(this,f,e))},resize:function(){var e=c("[data-dropdown-content].open"),f=c("[data-dropdown='"+e.attr("id")+"']");e.length&&f.length&&this.css(e,f)},css:function(j,g){var e=j.offsetParent();if(e.length>0&&/body/i.test(j.offsetParent()[0].nodeName)){var f=g.offset();f.top-=j.offsetParent().offset().top,f.left-=j.offsetParent().offset().left}else{var f=g.position()}if(this.small()){j.css({position:"absolute",width:"95%",left:"2.5%","max-width":"none",top:f.top+this.outerHeight(g)})}else{if(!Foundation.rtl&&c(a).width()>this.outerWidth(j)+g.offset().left){var h=f.left;j.hasClass("right")&&j.removeClass("right")}else{j.hasClass("right")||j.addClass("right");var h=f.left-(this.outerWidth(j)-this.outerWidth(g))}j.attr("style","").css({position:"absolute",top:f.top+this.outerHeight(g),left:h})}return j},small:function(){return c(a).width()<768||c("html").hasClass("lt-ie9")},off:function(){c(this.scope).off(".fndtn.dropdown"),c("html, body").off(".fndtn.dropdown"),c(a).off(".fndtn.dropdown"),c("[data-dropdown-content]").off(".fndtn.dropdown"),this.settings.init=!1},reflow:function(){}}}(Foundation.zj,this,this.document),function(m,w,g){function k(f){var a={},c=/^jQuery\d+$/;return g.each(f.attributes,function(i,l){l.specified&&!c.test(l.name)&&(a[l.name]=l.value)}),a}function h(l,f){var a=this,c=g(a);if(a.value==c.attr("placeholder")&&c.hasClass("placeholder")){if(c.data("placeholder-password")){c=c.hide().next().show().attr("id",c.removeAttr("id").data("placeholder-id"));if(l===!0){return c[0].value=f}c.focus()}else{a.value="",c.removeClass("placeholder"),a==w.activeElement&&a.select()}}}function p(){var n,c=this,l=g(c),a=l,f=this.id;if(c.value==""){if(c.type=="password"){if(!l.data("placeholder-textinput")){try{n=l.clone().attr({type:"text"})}catch(u){n=g("<input>").attr(g.extend(k(this),{type:"text"}))}n.removeAttr("name").data({"placeholder-password":!0,"placeholder-id":f}).bind("focus.placeholder",h),l.data({"placeholder-textinput":n,"placeholder-id":f}).before(n)}l=l.removeAttr("id").hide().prev().attr("id",f).show()}l.addClass("placeholder"),l[0].value=l.attr("placeholder")}else{l.removeClass("placeholder")}}var b="placeholder" in w.createElement("input"),j="placeholder" in w.createElement("textarea"),x=g.fn,d=g.valHooks,v,q;b&&j?(q=x.placeholder=function(){return this},q.input=q.textarea=!0):(q=x.placeholder=function(){var a=this;return a.filter((b?"textarea":":input")+"[placeholder]").not(".placeholder").bind({"focus.placeholder":h,"blur.placeholder":p}).data("placeholder-enabled",!0).trigger("blur.placeholder"),a},q.input=b,q.textarea=j,v={get:function(c){var a=g(c);return a.data("placeholder-enabled")&&a.hasClass("placeholder")?"":c.value},set:function(f,c){var a=g(f);return a.data("placeholder-enabled")?(c==""?(f.value=c,f!=w.activeElement&&p.call(f)):a.hasClass("placeholder")?h.call(f,!0,c)||(f.value=c):f.value=c,a):f.value=c}},b||(d.input=v),j||(d.textarea=v),g(function(){g(w).delegate("form","submit.placeholder",function(){var a=g(".placeholder",this).each(h);setTimeout(function(){a.each(p)},10)})}),g(m).bind("beforeunload.placeholder",function(){g(".placeholder").each(function(){this.value=""})}))}(this,document,Foundation.zj),function(d,b,f,c){Foundation.libs.forms={name:"forms",version:"4.2.1",cache:{},settings:{disable_class:"no-custom",last_combo:null},init:function(e,h,g){return typeof h=="object"&&d.extend(!0,this.settings,h),typeof h!="string"?(this.settings.init||this.events(),this.assemble(),this.settings.init):this[h].call(this,g)},assemble:function(){d('form.custom input[type="radio"]',d(this.scope)).not('[data-customforms="disabled"]').each(this.append_custom_markup),d('form.custom input[type="checkbox"]',d(this.scope)).not('[data-customforms="disabled"]').each(this.append_custom_markup),d("form.custom select",d(this.scope)).not('[data-customforms="disabled"]').not("[multiple=multiple]").each(this.append_custom_select)},events:function(){var e=this;d(this.scope).on("click.fndtn.forms","form.custom span.custom.checkbox",function(g){g.preventDefault(),g.stopPropagation(),e.toggle_checkbox(d(this))}).on("click.fndtn.forms","form.custom span.custom.radio",function(g){g.preventDefault(),g.stopPropagation(),e.toggle_radio(d(this))}).on("change.fndtn.forms",'form.custom select:not([data-customforms="disabled"])',function(g,h){e.refresh_custom_select(d(this),h)}).on("click.fndtn.forms","form.custom label",function(h){if(d(h.target).is("label")){var k=d("#"+e.escape(d(this).attr("for"))+':not([data-customforms="disabled"])'),g,j;k.length!==0&&(k.attr("type")==="checkbox"?(h.preventDefault(),g=d(this).find("span.custom.checkbox"),g.length==0&&(g=k.add(this).siblings("span.custom.checkbox").first()),e.toggle_checkbox(g)):k.attr("type")==="radio"&&(h.preventDefault(),j=d(this).find("span.custom.radio"),j.length==0&&(j=k.add(this).siblings("span.custom.radio").first()),e.toggle_radio(j)))}}).on("mousedown.fndtn.forms","form.custom div.custom.dropdown",function(){return !1}).on("click.fndtn.forms","form.custom div.custom.dropdown a.current, form.custom div.custom.dropdown a.selector",function(g){var j=d(this),h=j.closest("div.custom.dropdown"),i=a(h,"select");h.hasClass("open")||d(e.scope).trigger("click"),g.preventDefault();if(!1===i.is(":disabled")){return h.toggleClass("open"),h.hasClass("open")?d(e.scope).on("click.fndtn.forms.customdropdown",function(){h.removeClass("open"),d(e.scope).off(".fndtn.forms.customdropdown")}):d(e.scope).on(".fndtn.forms.customdropdown"),!1}}).on("click.fndtn.forms touchend.fndtn.forms","form.custom div.custom.dropdown li",function(h){var l=d(this),j=l.closest("div.custom.dropdown"),i=a(j,"select"),k=0;h.preventDefault(),h.stopPropagation();if(!d(this).hasClass("disabled")){d("div.dropdown").not(j).removeClass("open");var g=l.closest("ul").find("li.selected");g.removeClass("selected"),l.addClass("selected"),j.removeClass("open").find("a.current").text(l.text()),l.closest("ul").find("li").each(function(m){l[0]==this&&(k=m)}),i[0].selectedIndex=k,i.data("prevalue",g.html()),i.trigger("change")}}),d(b).on("keydown",function(k){var m=f.activeElement,j=Foundation.libs.forms,l=d(".custom.dropdown.open");if(l.length>0){k.preventDefault(),k.which===13&&l.find("li.selected").trigger("click"),k.which===27&&l.removeClass("open");if(k.which>=65&&k.which<=90){var n=j.go_to(l,k.which),h=l.find("li.selected");n&&(h.removeClass("selected"),j.scrollTo(n.addClass("selected"),300))}if(k.which===38){var h=l.find("li.selected"),g=h.prev(":not(.disabled)");g.length>0&&(g.parent()[0].scrollTop=g.parent().scrollTop()-j.outerHeight(g),h.removeClass("selected"),g.addClass("selected"))}else{if(k.which===40){var h=l.find("li.selected"),n=h.next(":not(.disabled)");n.length>0&&(n.parent()[0].scrollTop=n.parent().scrollTop()+j.outerHeight(n),h.removeClass("selected"),n.addClass("selected"))}}}}),this.settings.init=!0},go_to:function(l,h){var m=l.find("li"),k=m.length;if(k>0){for(var g=0;g<k;g++){var j=m.eq(g).text().charAt(0).toLowerCase();if(j===String.fromCharCode(h).toLowerCase()){return m.eq(g)}}}},scrollTo:function(l,h){if(h<0){return}var m=l.parent(),k=this.outerHeight(l),g=k*l.index()-m.scrollTop(),j=g/h*10;this.scrollToTimerCache=setTimeout(function(){isNaN(parseInt(j,10))||(m[0].scrollTop=m.scrollTop()+j,this.scrollTo(l,h-10))}.bind(this),10)},append_custom_markup:function(g,k){var j=d(k),e=j.attr("type"),h=j.next("span.custom."+e);h.length===0&&(h=d('<span class="custom '+e+'"></span>').insertAfter(j)),h.toggleClass("checked",j.is(":checked")),h.toggleClass("disabled",j.is(":disabled"))},append_custom_select:function(F,k){var e=Foundation.libs.forms,x=d(k),G=x.next("div.custom.dropdown"),j=G.find("ul"),E=G.find(".current"),C=G.find(".selector"),z=x.find("option"),w=z.filter(":selected"),B=x.attr("class")?x.attr("class").split(" "):[],y=0,g="",A,D=!1;if(x.hasClass(e.settings.disable_class)){return}if(G.length===0){var q=x.hasClass("small")?"small":x.hasClass("medium")?"medium":x.hasClass("large")?"large":x.hasClass("expand")?"expand":"";G=d('<div class="'+["custom","dropdown",q].concat(B).filter(function(i,h,l){return i==""?!1:l.indexOf(i)==h}).join(" ")+'"><a href="#" class="selector"></a><ul /></div>'),C=G.find(".selector"),j=G.find("ul"),g=z.map(function(){var h=d(this).attr("class")?d(this).attr("class"):"";return"<li class='"+h+"'>"+d(this).html()+"</li>"}).get().join(""),j.append(g),D=G.prepend('<a href="#" class="current">'+w.html()+"</a>").find(".current"),x.after(G).addClass("hidden-field")}else{g=z.map(function(){return"<li>"+d(this).html()+"</li>"}).get().join(""),j.html("").append(g)}e.assign_id(x,G),G.toggleClass("disabled",x.is(":disabled")),A=j.find("li"),e.cache[G.data("id")]=A.length,z.each(function(h){this.selected&&(A.eq(h).addClass("selected"),D&&D.html(d(this).html())),d(this).is(":disabled")&&A.eq(h).addClass("disabled")});if(!G.is(".small, .medium, .large, .expand")){G.addClass("open");var e=Foundation.libs.forms;e.hidden_fix.adjust(j),y=e.outerWidth(A)>y?e.outerWidth(A):y,Foundation.libs.forms.hidden_fix.reset(),G.removeClass("open")}},assign_id:function(h,g){var i=[+(new Date),Foundation.random_str(5)].join("-");h.attr("data-id",i),g.attr("data-id",i)},refresh_custom_select:function(h,m){var k=this,g=0,j=h.next(),l=h.find("option"),e=j.find("li");if(e.length!=this.cache[j.data("id")]||m){j.find("ul").html(""),l.each(function(){var i=d("<li>"+d(this).html()+"</li>");j.find("ul").append(i)}),l.each(function(i){this.selected&&(j.find("li").eq(i).addClass("selected"),j.find(".current").html(d(this).html())),d(this).is(":disabled")&&j.find("li").eq(i).addClass("disabled")}),j.removeAttr("style").find("ul").removeAttr("style"),j.find("li").each(function(){j.addClass("open"),k.outerWidth(d(this))>g&&(g=k.outerWidth(d(this))),j.removeClass("open")}),e=j.find("li"),this.cache[j.data("id")]=e.length}},toggle_checkbox:function(h){var g=h.prev(),i=g[0];!1===g.is(":disabled")&&(i.checked=i.checked?!1:!0,h.toggleClass("checked"),g.trigger("change"))},toggle_radio:function(i){var g=i.prev(),j=g.closest("form.custom"),h=g[0];!1===g.is(":disabled")&&(j.find('input[type="radio"][name="'+this.escape(g.attr("name"))+'"]').next().not(i).removeClass("checked"),i.hasClass("checked")||i.toggleClass("checked"),h.checked=i.hasClass("checked"),g.trigger("change"))},escape:function(g){return g.replace(/[-[\]{}()*+?.,\\^$|#\s]/g,"\\$&")},hidden_fix:{tmp:[],hidden:null,adjust:function(e){var g=this;g.hidden=e.parents(),g.hidden=g.hidden.add(e).filter(":hidden"),g.hidden.each(function(){var h=d(this);g.tmp.push(h.attr("style")),h.css({visibility:"hidden",display:"block"})})},reset:function(){var e=this;e.hidden.each(function(j){var g=d(this),h=e.tmp[j];h===c?g.removeAttr("style"):g.attr("style",h)}),e.tmp=[],e.hidden=null}},off:function(){d(this.scope).off(".fndtn.forms")},reflow:function(){}};var a=function(e,g){var e=e.prev();while(e.length){if(e.is(g)){return e}e=e.prev()}return d()}}(Foundation.zj,this,this.document),function(c,a,d,b){Foundation.libs.section={name:"section",version:"4.2.0",settings:{deep_linking:!1,one_up:!0,section_selector:"[data-section]",region_selector:"section, .section, [data-section-region]",title_selector:".title, [data-section-title]",active_region_selector:"section.active, .section.active, .active[data-section-region]",content_selector:".content, [data-section-content]",nav_selector:'[data-section="vertical-nav"], [data-section="horizontal-nav"]',callback:function(){}},init:function(f,h,g){var e=this;return Foundation.inherit(this,"throttle data_options position_right offset_right"),typeof h=="object"&&c.extend(!0,e.settings,h),typeof h!="string"?(this.set_active_from_hash(),this.events(),!0):this[h].call(this,g)},events:function(){var e=this;c(this.scope).on("click.fndtn.section","[data-section] .title, [data-section] [data-section-title]",function(g){var h=c(this),f=h.closest(e.settings.region_selector);f.children(e.settings.content_selector).length>0&&(e.toggle_active.call(this,g,e),e.reflow())}),c(a).on("resize.fndtn.section",e.throttle(function(){e.resize.call(this)},30)).on("hashchange",function(){e.settings.toggled||(e.set_active_from_hash(),c(this).trigger("resize"))}).trigger("resize"),c(d).on("click.fndtn.section",function(f){c(f.target).closest(e.settings.title_selector).length<1&&c(e.settings.nav_selector).children(e.settings.region_selector).removeClass("active").attr("style","")})},toggle_active:function(p,h){var e=c(this),h=Foundation.libs.section,j=e.closest(h.settings.region_selector),q=e.siblings(h.settings.content_selector),g=j.parent(),m=c.extend({},h.settings,h.data_options(g)),l=g.children(h.settings.active_region_selector);h.settings.toggled=!0,!m.deep_linking&&q.length>0&&p.preventDefault();if(j.hasClass("active")){(h.small(g)||h.is_vertical_nav(g)||h.is_horizontal_nav(g)||h.is_accordion(g))&&(l[0]!==j[0]||l[0]===j[0]&&!m.one_up)&&j.removeClass("active").attr("style","")}else{var l=g.children(h.settings.active_region_selector),k=h.outerHeight(j.children(h.settings.title_selector));if(h.small(g)||m.one_up){h.small(g)?l.attr("style",""):l.attr("style","visibility: hidden; padding-top: "+k+"px;")}h.small(g)?j.attr("style",""):j.css("padding-top",k),j.addClass("active"),l.length>0&&l.removeClass("active").attr("style",""),h.is_vertical_tabs(g)&&(q.css("display","block"),l!==null&&l.children(h.settings.content_selector).css("display","none"))}setTimeout(function(){h.settings.toggled=!1},300),m.callback()},resize:function(){var e=Foundation.libs.section,f=c(e.settings.section_selector);f.each(function(){var k=c(this),j=k.children(e.settings.active_region_selector),g=c.extend({},e.settings,e.data_options(k));if(j.length>1){j.not(":first").removeClass("active").attr("style","")}else{if(j.length<1&&!e.is_vertical_nav(k)&&!e.is_horizontal_nav(k)&&!e.is_accordion(k)){var h=k.children(e.settings.region_selector).first();(g.one_up||!e.small(k))&&h.addClass("active"),e.small(k)?h.attr("style",""):h.css("padding-top",e.outerHeight(h.children(e.settings.title_selector)))}}e.small(k)?j.attr("style",""):j.css("padding-top",e.outerHeight(j.children(e.settings.title_selector))),e.position_titles(k),e.is_horizontal_nav(k)&&!e.small(k)||e.is_vertical_tabs(k)&&!e.small(k)?e.position_content(k):e.position_content(k,!1)})},is_vertical_nav:function(f){return/vertical-nav/i.test(f.data("section"))},is_horizontal_nav:function(f){return/horizontal-nav/i.test(f.data("section"))},is_accordion:function(f){return/accordion/i.test(f.data("section"))},is_horizontal_tabs:function(f){return/^tabs$/i.test(f.data("section"))},is_vertical_tabs:function(f){return/vertical-tabs/i.test(f.data("section"))},set_active_from_hash:function(){var g=a.location.hash.substring(1),f=c("[data-section]"),e=this;f.each(function(){var j=c(this),l=c.extend({},e.settings,e.data_options(j));if(g.length>0&&l.deep_linking){var k=j.children(e.settings.region_selector).attr("style","").removeClass("active"),m=k.map(function(){var n=c(e.settings.content_selector,this),p=n.data("slug");if((new RegExp(p,"i")).test(g)){return n}}),i=m.length;for(var h=i-1;h>=0;h--){c(m[h]).parent().addClass("active")}}})},position_titles:function(f,k){var h=this,e=f.children(this.settings.region_selector).map(function(){return c(this).children(h.settings.title_selector)}),g=0,j=0,h=this;typeof k=="boolean"?e.attr("style",""):e.each(function(){h.is_vertical_tabs(f)?(c(this).css("top",j),j+=h.outerHeight(c(this))):(h.rtl?c(this).css("right",g):c(this).css("left",g),g+=h.outerWidth(c(this)))})},position_content:function(p,h){var e=this,j=p.children(e.settings.region_selector),q=j.map(function(){return c(this).children(e.settings.title_selector)}),g=j.map(function(){return c(this).children(e.settings.content_selector)});if(typeof h=="boolean"){g.attr("style",""),p.attr("style","")}else{if(e.is_vertical_tabs(p)&&!e.small(p)){var m=0,l=Number.MAX_VALUE,k=null;j.each(function(){var u=c(this),f=u.children(e.settings.title_selector),r=u.children(e.settings.content_selector),t=0;k=e.outerWidth(f),t=e.outerWidth(p)-k,t<l&&(l=t),m+=e.outerHeight(f),c(this).hasClass("active")||r.css("display","none")}),j.each(function(){var f=c(this).children(e.settings.content_selector);f.css("minHeight",m),f.css("maxWidth",l-2)})}else{j.each(function(){var r=c(this),s=r.children(e.settings.title_selector),f=r.children(e.settings.content_selector);e.rtl?f.css({right:e.position_right(s)+1,top:e.outerHeight(s)-2}):f.css({left:s.position().left-1,top:e.outerHeight(s)-2})}),typeof Zepto=="function"?p.height(this.outerHeight(c(q[0]))):p.height(this.outerHeight(c(q[0]))-2)}}},position_right:function(f){var k=this,h=f.closest(this.settings.section_selector),e=h.children(this.settings.region_selector),g=f.closest(this.settings.section_selector).width(),j=e.map(function(){return c(this).children(k.settings.title_selector)}).length;return g-f.position().left-f.width()*(f.index()+1)-j},reflow:function(e){var e=e||d;c(this.settings.section_selector,e).trigger("resize")},small:function(e){var f=c.extend({},this.settings,this.data_options(e));return this.is_horizontal_tabs(e)?!1:e&&this.is_accordion(e)?!0:c("html").hasClass("lt-ie9")?!0:c("html").hasClass("ie8compat")?!0:c(this.scope).width()<768},off:function(){c(this.scope).off(".fndtn.section"),c(a).off(".fndtn.section"),c(d).off(".fndtn.section")}}}(Foundation.zj,this,this.document),function(c,a,d,b){Foundation.libs.topbar={name:"topbar",version:"4.2.0",settings:{index:0,stickyClass:"sticky",custom_back_text:!0,back_text:"Back",is_hover:!0,scrolltop:!0,init:!1},init:function(h,g,e){Foundation.inherit(this,"data_options");var f=this;return typeof g=="object"?c.extend(!0,this.settings,g):typeof e!="undefined"&&c.extend(!0,this.settings,e),typeof g!="string"?(c(".top-bar, [data-topbar]").each(function(){c.extend(!0,f.settings,f.data_options(c(this))),f.settings.$w=c(a),f.settings.$topbar=c(this),f.settings.$section=f.settings.$topbar.find("section"),f.settings.$titlebar=f.settings.$topbar.children("ul").first(),f.settings.$topbar.data("index",0);var i=c("<div class='top-bar-js-breakpoint'/>").insertAfter(f.settings.$topbar);f.settings.breakPoint=i.width(),i.remove(),f.assemble(),f.settings.$topbar.parent().hasClass("fixed")&&c("body").css("padding-top",f.outerHeight(f.settings.$topbar))}),f.settings.init||this.events(),this.settings.init):this[g].call(this,e)},events:function(){var f=this,e=this.outerHeight(c(".top-bar, [data-topbar]"));c(this.scope).off(".fndtn.topbar").on("click.fndtn.topbar",".top-bar .toggle-topbar, [data-topbar] .toggle-topbar",function(h){var j=c(this).closest(".top-bar, [data-topbar]"),k=j.find("section, .section"),g=j.children("ul").first();h.preventDefault(),f.breakpoint()&&(f.rtl?(k.css({right:"0%"}),k.find(">.name").css({right:"100%"})):(k.css({left:"0%"}),k.find(">.name").css({left:"100%"})),k.find("li.moved").removeClass("moved"),j.data("index",0),j.toggleClass("expanded").css("min-height","")),j.hasClass("expanded")?j.parent().hasClass("fixed")&&(j.parent().removeClass("fixed"),j.addClass("fixed"),c("body").css("padding-top","0"),f.settings.scrolltop&&a.scrollTo(0,0)):j.hasClass("fixed")&&(j.parent().addClass("fixed"),j.removeClass("fixed"),c("body").css("padding-top",e))}).on("mouseenter mouseleave",".top-bar li",function(g){if(!f.settings.is_hover){return}/enter|over/i.test(g.type)?c(this).addClass("hover"):c(this).removeClass("hover")}).on("click.fndtn.topbar",".top-bar li.has-dropdown",function(h){if(f.breakpoint()){return}var k=c(this),g=c(h.target),j=k.closest("[data-topbar], .top-bar"),l=j.data("topbar");if(f.settings.is_hover&&!Modernizr.touch){return}h.stopImmediatePropagation(),g[0].nodeName==="A"&&g.parent().hasClass("has-dropdown")&&h.preventDefault(),k.hasClass("hover")?k.removeClass("hover").find("li").removeClass("hover"):k.addClass("hover")}).on("click.fndtn.topbar",".top-bar .has-dropdown>a, [data-topbar] .has-dropdown>a",function(k){if(f.breakpoint()){k.preventDefault();var m=c(this),j=m.closest(".top-bar, [data-topbar]"),l=j.find("section, .section"),n=j.children("ul").first(),h=m.next(".dropdown").outerHeight(),g=m.closest("li");j.data("index",j.data("index")+1),g.addClass("moved"),f.rtl?(l.css({right:-(100*j.data("index"))+"%"}),l.find(">.name").css({right:100*j.data("index")+"%"})):(l.css({left:-(100*j.data("index"))+"%"}),l.find(">.name").css({left:100*j.data("index")+"%"})),j.css("min-height",f.height(m.siblings("ul"))+f.outerHeight(n,!0))}}),c(a).on("resize.fndtn.topbar",function(){f.breakpoint()||c(".top-bar, [data-topbar]").css("min-height","").removeClass("expanded").find("li").removeClass("hover")}.bind(this)),c("body").on("click.fndtn.topbar",function(g){var h=c(g.target).closest("[data-topbar], .top-bar");if(h.length>0){return}c(".top-bar li, [data-topbar] li").removeClass("hover")}),c(this.scope).on("click.fndtn",".top-bar .has-dropdown .back, [data-topbar] .has-dropdown .back",function(k){k.preventDefault();var m=c(this),j=m.closest(".top-bar, [data-topbar]"),l=j.children("ul").first(),n=j.find("section, .section"),h=m.closest("li.moved"),g=h.parent();j.data("index",j.data("index")-1),f.rtl?(n.css({right:-(100*j.data("index"))+"%"}),n.find(">.name").css({right:100*j.data("index")+"%"})):(n.css({left:-(100*j.data("index"))+"%"}),n.find(">.name").css({left:100*j.data("index")+"%"})),j.data("index")===0?j.css("min-height",0):j.css("min-height",f.height(g)+f.outerHeight(l,!0)),setTimeout(function(){h.removeClass("moved")},300)})},breakpoint:function(){return c(a).width()<=this.settings.breakPoint||c("html").hasClass("lt-ie9")},assemble:function(){var e=this;this.settings.$section.detach(),this.settings.$section.find(".has-dropdown>a").each(function(){var j=c(this),h=j.siblings(".dropdown"),f=j.attr("href");if(f&&f.length>1){var g=c('<li class="title back js-generated"><h5><a href="#"></a></h5></li><li><a class="parent-link js-generated" href="'+f+'">'+j.text()+"</a></li>")}else{var g=c('<li class="title back js-generated"><h5><a href="#"></a></h5></li>')}e.settings.custom_back_text==1?g.find("h5>a").html("&laquo; "+e.settings.back_text):g.find("h5>a").html("&laquo; "+j.html()),h.prepend(g)}),this.settings.$section.appendTo(this.settings.$topbar),this.sticky()},height:function(e){var g=0,f=this;return e.find("> li").each(function(){g+=f.outerHeight(c(this),!0)}),g},sticky:function(){var h="."+this.settings.stickyClass;if(c(h).length>0){var g=c(h).length?c(h).offset().top:0,e=c(a),f=this.outerHeight(c(".top-bar"));e.scroll(function(){e.scrollTop()>=g?(c(h).addClass("fixed"),c("body").css("padding-top",f)):e.scrollTop()<g&&(c(h).removeClass("fixed"),c("body").css("padding-top","0"))})}},off:function(){c(this.scope).off(".fndtn.topbar"),c(a).off(".fndtn.topbar")},reflow:function(){}}}(Foundation.zj,this,this.document);$(function(){$(document).foundation();var d;var h=$("#msg");var b=$("#log");var c=$("#form");function f(k){var j=b[0];var i=j.scrollTop==j.scrollHeight-j.clientHeight;k.appendTo(b);if(i){j.scrollTop=j.scrollHeight-j.clientHeight}}function e(){return(new Date().toTimeString().replace(/.*(\d{2}:\d{2}:\d{2}).*/,"$1"))}function a(j){var i=j.Arguments.Nickname;var k;switch(j.Event){case 0:k=" <strong>"+i+":</strong> "+j.Arguments.Body;break;case 1:k=" <strong>*** "+i+" has joined ***</strong> ";g(j.Arguments.Stats);break;case 2:k=" <strong>*** "+i+" has left ***</strong> ";g(j.Arguments.Stats);break}f($("<div class='row'><div class='large-12 columns msg'> <p><strong><span class='timestamp'>"+e()+"</span></strong>"+k+"</p> </div></div>"))}function g(i){$("#room-stats").text(i.UsersCount)}h.on("keyup",function(i){i=i||event;if(i.keyCode===13&&!i.shiftKey){c.submit()}return true});c.submit(function(){if(!d){return false}if(!h.val()){return false}d.send(h.val());h.val("");return false});if(window.WebSocket){d=new WebSocket(c.data("socket"));d.onclose=function(i){f($("<div class='row'><div class='large-12 columns msg'> <p><strong><span class='timestamp'>"+e()+"</span></strong>   <strong>*** CONNECTION CLOSED ***</strong></p> </div></div>"))};d.onmessage=function(i){o=jQuery.parseJSON(i.data);a(o)}}else{f($("<div class='row'><div class='large-12 columns msg'><p><strong>*** Your browser doesn't support WebSockets. ***</strong></p> </div></div>"))}});