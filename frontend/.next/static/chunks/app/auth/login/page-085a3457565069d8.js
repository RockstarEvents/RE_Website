(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[716],{7467:function(e,r,t){Promise.resolve().then(t.bind(t,4167))},8377:function(e,r,t){"use strict";t.d(r,{U:function(){return h}});var s=t(7437),n=t(2265),a=t(7440);let i=n.forwardRef((e,r)=>{let{className:t,...n}=e;return(0,s.jsx)("div",{ref:r,className:(0,a.cn)("rounded-xl border bg-card text-card-foreground shadow",t),...n})});i.displayName="Card";let l=n.forwardRef((e,r)=>{let{className:t,...n}=e;return(0,s.jsx)("div",{ref:r,className:(0,a.cn)("flex flex-col space-y-1.5 p-6",t),...n})});l.displayName="CardHeader",n.forwardRef((e,r)=>{let{className:t,...n}=e;return(0,s.jsx)("h3",{ref:r,className:(0,a.cn)("font-semibold leading-none tracking-tight",t),...n})}).displayName="CardTitle",n.forwardRef((e,r)=>{let{className:t,...n}=e;return(0,s.jsx)("p",{ref:r,className:(0,a.cn)("text-sm text-muted-foreground",t),...n})}).displayName="CardDescription";let d=n.forwardRef((e,r)=>{let{className:t,...n}=e;return(0,s.jsx)("div",{ref:r,className:(0,a.cn)("p-6 pt-0",t),...n})});d.displayName="CardContent";let o=n.forwardRef((e,r)=>{let{className:t,...n}=e;return(0,s.jsx)("div",{ref:r,className:(0,a.cn)("flex items-center p-6 pt-0",t),...n})});o.displayName="CardFooter";var c=t(2448),u=t.n(c);let m=e=>{let{label:r}=e;return(0,s.jsxs)("div",{className:"w-full flex flex-col gap-y-4 items-center justify-center",children:[(0,s.jsx)("h1",{className:(0,a.cn)("text-3xl font-semibold",u().className),children:"Ramos-project"}),(0,s.jsx)("p",{className:"text-sm",children:r})]})};var f=t(495),x=t(7138);let p=e=>{let{href:r,label:t}=e;return(0,s.jsx)(f.z,{variant:"link",className:"font-normal w-full",size:"sm",asChild:!0,children:(0,s.jsx)(x.default,{href:r,children:t})})},h=e=>{let{children:r,headerLabel:t,backButtonLabel:n,backButtonHref:a}=e;return(0,s.jsxs)(i,{className:"w-[400px] shadow-md",children:[(0,s.jsx)(l,{children:(0,s.jsx)(m,{label:t})}),(0,s.jsx)(d,{children:r}),(0,s.jsx)(o,{children:(0,s.jsx)(p,{label:n,href:a})})]})}},4167:function(e,r,t){"use strict";t.d(r,{LoginForm:function(){return h}});var s=t(7437),n=t(2265),a=t(9343),i=t(1014),l=t(6661),d=t(4609),o=t(3102),c=t(8377),u=t(495),m=t(4867);let f=e=>{let{message:r}=e;return r?(0,s.jsxs)("div",{className:"bg-destructive/15 p-3 rounded-md flex items-center gap-x-2 text-sm text-destructive",children:[(0,s.jsx)(m.LPM,{className:"h-4 w-4"}),(0,s.jsx)("p",{children:r})]}):null},x=e=>{let{message:r}=e;return r?(0,s.jsxs)("div",{className:"bg-emerald-500/15 p-3 rounded-md flex items-center gap-x-2 text-sm text-emerald-500",children:[(0,s.jsx)(m.NhS,{className:"h-4 w-4"}),(0,s.jsx)("p",{children:r})]}):null};t(4590);var p=(0,t(8064).$)("ff6d23ce888597eeb09915b6f321cd2371633cb2");let h=()=>{let[e,r]=(0,n.useState)(""),[t,m]=(0,n.useState)(""),[h,b]=(0,n.useTransition)(),g=(0,a.cI)({resolver:(0,i.F)(d.p),defaultValues:{email:"",password:""}});return(0,s.jsx)(c.U,{headerLabel:"Welcome",backButtonLabel:"Register",backButtonHref:"/auth/register",children:(0,s.jsx)(l.l0,{...g,children:(0,s.jsxs)("form",{onSubmit:g.handleSubmit(e=>{r(""),m(""),b(()=>{p(e).then(e=>{r(e.error),m(e.succes)})})}),className:"space-y-8",children:[(0,s.jsxs)("div",{className:"space-y-5",children:[(0,s.jsx)(l.Wi,{control:g.control,name:"email",render:e=>{let{field:r}=e;return(0,s.jsxs)(l.xJ,{children:[(0,s.jsx)(l.lX,{children:"Email"}),(0,s.jsx)(l.NI,{children:(0,s.jsx)(o.I,{...r,disabled:h,placeholder:"sergio.ramos@mail.ru",type:"email"})}),(0,s.jsx)(l.zG,{})]})}}),(0,s.jsx)(l.Wi,{control:g.control,name:"password",render:e=>{let{field:r}=e;return(0,s.jsxs)(l.xJ,{children:[(0,s.jsx)(l.lX,{children:"Password"}),(0,s.jsx)(l.NI,{children:(0,s.jsx)(o.I,{...r,disabled:h,placeholder:"******",type:"password"})}),(0,s.jsx)(l.zG,{})]})}})]}),(0,s.jsx)(f,{message:e}),(0,s.jsx)(x,{message:t}),(0,s.jsx)(u.z,{disabled:h,type:"submit",className:"w-full",children:"Login"})]})})})}},495:function(e,r,t){"use strict";t.d(r,{z:function(){return o}});var s=t(7437),n=t(2265),a=t(2974),i=t(2218),l=t(7440);let d=(0,i.j)("inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50",{variants:{variant:{default:"bg-primary text-primary-foreground shadow hover:bg-primary/90",destructive:"bg-destructive text-destructive-foreground shadow-sm hover:bg-destructive/90",outline:"border border-input bg-background shadow-sm hover:bg-accent hover:text-accent-foreground",secondary:"bg-secondary text-secondary-foreground shadow-sm hover:bg-secondary/80",ghost:"hover:bg-accent hover:text-accent-foreground",link:"text-primary underline-offset-4 hover:underline"},size:{default:"h-9 px-4 py-2",sm:"h-8 rounded-md px-3 text-xs",lg:"h-10 rounded-md px-8",icon:"h-9 w-9"}},defaultVariants:{variant:"default",size:"default"}}),o=n.forwardRef((e,r)=>{let{className:t,variant:n,size:i,asChild:o=!1,...c}=e,u=o?a.g7:"button";return(0,s.jsx)(u,{className:(0,l.cn)(d({variant:n,size:i,className:t})),ref:r,...c})});o.displayName="Button"},6661:function(e,r,t){"use strict";t.d(r,{l0:function(){return u},NI:function(){return g},Wi:function(){return f},xJ:function(){return h},lX:function(){return b},zG:function(){return j}});var s=t(7437),n=t(2265),a=t(2974),i=t(9343),l=t(7440),d=t(7200);let o=(0,t(2218).j)("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"),c=n.forwardRef((e,r)=>{let{className:t,...n}=e;return(0,s.jsx)(d.f,{ref:r,className:(0,l.cn)(o(),t),...n})});c.displayName=d.f.displayName;let u=i.RV,m=n.createContext({}),f=e=>{let{...r}=e;return(0,s.jsx)(m.Provider,{value:{name:r.name},children:(0,s.jsx)(i.Qr,{...r})})},x=()=>{let e=n.useContext(m),r=n.useContext(p),{getFieldState:t,formState:s}=(0,i.Gc)(),a=t(e.name,s);if(!e)throw Error("useFormField should be used within <FormField>");let{id:l}=r;return{id:l,name:e.name,formItemId:"".concat(l,"-form-item"),formDescriptionId:"".concat(l,"-form-item-description"),formMessageId:"".concat(l,"-form-item-message"),...a}},p=n.createContext({}),h=n.forwardRef((e,r)=>{let{className:t,...a}=e,i=n.useId();return(0,s.jsx)(p.Provider,{value:{id:i},children:(0,s.jsx)("div",{ref:r,className:(0,l.cn)("space-y-2",t),...a})})});h.displayName="FormItem";let b=n.forwardRef((e,r)=>{let{className:t,...n}=e,{error:a,formItemId:i}=x();return(0,s.jsx)(c,{ref:r,className:(0,l.cn)(a&&"text-destructive",t),htmlFor:i,...n})});b.displayName="FormLabel";let g=n.forwardRef((e,r)=>{let{...t}=e,{error:n,formItemId:i,formDescriptionId:l,formMessageId:d}=x();return(0,s.jsx)(a.g7,{ref:r,id:i,"aria-describedby":n?"".concat(l," ").concat(d):"".concat(l),"aria-invalid":!!n,...t})});g.displayName="FormControl",n.forwardRef((e,r)=>{let{className:t,...n}=e,{formDescriptionId:a}=x();return(0,s.jsx)("p",{ref:r,id:a,className:(0,l.cn)("text-[0.8rem] text-muted-foreground",t),...n})}).displayName="FormDescription";let j=n.forwardRef((e,r)=>{let{className:t,children:n,...a}=e,{error:i,formMessageId:d}=x(),o=i?String(null==i?void 0:i.message):n;return o?(0,s.jsx)("p",{ref:r,id:d,className:(0,l.cn)("text-[0.8rem] font-medium text-destructive",t),...a,children:o}):null});j.displayName="FormMessage"},3102:function(e,r,t){"use strict";t.d(r,{I:function(){return i}});var s=t(7437),n=t(2265),a=t(7440);let i=n.forwardRef((e,r)=>{let{className:t,type:n,...i}=e;return(0,s.jsx)("input",{type:n,className:(0,a.cn)("flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50",t),ref:r,...i})});i.displayName="Input"},7440:function(e,r,t){"use strict";t.d(r,{cn:function(){return a}});var s=t(4839),n=t(6164);function a(){for(var e=arguments.length,r=Array(e),t=0;t<e;t++)r[t]=arguments[t];return(0,n.m6)((0,s.W)(r))}},8064:function(e,r,t){"use strict";Object.defineProperty(r,"$",{enumerable:!0,get:function(){return n}});let s=t(4590);function n(e){let{createServerReference:r}=t(6671);return r(e,s.callServer)}},4609:function(e,r,t){"use strict";t.d(r,{T:function(){return a},p:function(){return n}});var s=t(9772);let n=s.Ry({email:s.Z_().email({message:"Email is required"}),password:s.Z_().min(1,{message:"Password is required"})}),a=s.Ry({name:s.Z_().min(1,{message:"Name is required"}),email:s.Z_().email({message:"Email is required"}),password:s.Z_().min(6,{message:"Minimum 6 characters required"})})}},function(e){e.O(0,[95,310,448,971,23,744],function(){return e(e.s=7467)}),_N_E=e.O()}]);