package i18n

// HTML strings — see en.go for tag and escaping rules.
//
// RTL strategy:
//   - Every FA message body is wrapped in <blockquote>...</blockquote>. Telegram
//     renders blockquote as a left-bordered side block that respects the
//     paragraph's natural BiDi direction; in our case the first strong char is
//     Persian, so the whole quote renders right-aligned.
//   - First visible character of every block is a Persian letter, never an
//     emoji or punctuation, so BiDi resolves to RTL even after editMessageText.
//   - U+200F (RLM) is kept as defense-in-depth at the very start.
//   - Inline button labels are NOT wrapped in blockquote (Telegram strips HTML
//     in button text); they stay short with the emoji at the end of the label.
const rlm = "\u200F"

var fa = map[Key]string{
	Welcome: rlm + "<blockquote><b>به ربات %s خوش آمدید 👋</b>\n\n" +
		"از این ربات می‌توانید:\n" +
		"• یک اکانت VPN موجود را لینک کنید 🔗\n" +
		"• مصرف و تاریخ انقضا را ببینید 📊\n" +
		"• درخواست تمدید بدهید 🔄\n" +
		"• سفارش اکانت جدید ثبت کنید 🆕\n" +
		"• هشدار خودکار کم‌شدن حجم بگیرید 🔔</blockquote>",

	BotDisabled: rlm + "<blockquote>ربات در حال حاضر توسط ادمین غیرفعال شده است. لطفاً بعداً مراجعه کنید. ⚠️</blockquote>",
	MainMenu:    rlm + "<blockquote><b>منوی اصلی 📋</b>\nچه کاری می‌خواهید انجام دهید؟</blockquote>",

	BtnAddAccount: rlm + "افزودن اکانت 🔗",
	BtnMyAccounts: rlm + "اکانت‌های من 👤",
	BtnNewOrder:   rlm + "سفارش اکانت جدید 🆕",
	BtnHelp:       rlm + "راهنما ℹ️",
	BtnLanguage:   rlm + "زبان 🌐",
	BtnCancel:     rlm + "انصراف ❌",
	BtnBack:       rlm + "بازگشت ⬅️",
	BtnUsage:      rlm + "مصرف 📊",
	BtnRenew:      rlm + "تمدید 🔄",
	BtnRemove:     rlm + "حذف لینک 🗑",

	AskUsername:    rlm + "<blockquote>لطفاً <b>نام کاربری</b> اکانت VPN خود را ارسال کنید: 🔑</blockquote>",
	AskPassword:    rlm + "<blockquote>حالا <b>رمز عبور</b> اکانت VPN را ارسال کنید (پیام شما بلافاصله حذف می‌شود): 🔒</blockquote>",
	AskUsernameNew: rlm + "<blockquote>یک نام کاربری برای اکانت جدید انتخاب کنید (۳ تا ۳۲ کاراکتر، حروف لاتین، عدد و <code>_ - .</code>): 📝</blockquote>",
	AskMessage:     rlm + "<blockquote>می‌توانید توضیح اختیاری برای ادمین بنویسید، یا /skip را بفرستید تا رد شوید: 💬</blockquote>",
	AskReceipt:     rlm + "<blockquote>لطفاً تصویر رسید پرداخت را به‌صورت <b>عکس</b> ارسال کنید. 🧾</blockquote>",

	AuthSuccess:   rlm + "<blockquote>اکانت با موفقیت لینک شد. ✅</blockquote>",
	AuthFail:      rlm + "<blockquote>نام کاربری یا رمز عبور اشتباه است. ❌</blockquote>",
	AuthLocked:    rlm + "<blockquote>اکانت شما قفل شده است. لطفاً با ادمین تماس بگیرید. 🔒</blockquote>",
	AlreadyLinked: rlm + "<blockquote>این اکانت قبلاً به چت تلگرام شما لینک شده است. ℹ️</blockquote>",

	NoAccounts:  rlm + "<blockquote>هنوز اکانتی لینک نکرده‌اید. از منوی اصلی <b>افزودن اکانت</b> را انتخاب کنید. 📭</blockquote>",
	NoPackages:  rlm + "<blockquote>در حال حاضر پکیج فعالی وجود ندارد. لطفاً بعداً مراجعه کنید. 📦</blockquote>",
	PickPackage: rlm + "<blockquote><b>یک پکیج انتخاب کنید: 📦</b></blockquote>",

	PickAccountRenew: rlm + "<blockquote>اکانتی را برای تمدید انتخاب کنید: 🔄</blockquote>",
	RequestCreated:   rlm + "<blockquote>درخواست شما ثبت شد. ادمین به‌زودی بررسی می‌کند. 📨</blockquote>",
	RequestExists:    rlm + "<blockquote>یک درخواست در حال بررسی دارید. لطفاً تا تعیین تکلیف صبر کنید. ⏳</blockquote>",
	WaitForApproval:  rlm + "<blockquote>در انتظار تایید ادمین… ⏳</blockquote>",
	NotApprovedYet:   rlm + "<blockquote>درخواست شما هنوز تایید نشده، امکان ارسال رسید وجود ندارد. ℹ️</blockquote>",
	ReceiptSaved:     rlm + "<blockquote>رسید دریافت شد. در انتظار تایید نهایی ادمین. 🧾</blockquote>",
	OnlyPhoto:        rlm + "<blockquote>لطفاً رسید را به‌صورت عکس ارسال کنید. 📷</blockquote>",

	HelpText: rlm + "<blockquote><b>راهنمای ربات</b>\n\n" +
		"از دکمه‌های inline برای مدیریت اکانت‌ها، مشاهدهٔ مصرف، درخواست تمدید، سفارش اکانت جدید و ارسال رسید پرداخت استفاده کنید.\n\n" +
		"دستورها:\n" +
		"• /start — منوی اصلی\n" +
		"• /help — این راهنما\n" +
		"• /settings — تنظیمات زبان\n" +
		"• /cancel — لغو عملیات</blockquote>",

	UsageText: rlm + "<blockquote><b>اکانت:</b> <code>%s</code> 👤\n" +
		"<b>وضعیت:</b> %s 📌\n" +
		"<b>حجم بسته:</b> %d GB 💾\n" +
		"<b>مصرف دریافت:</b> %.2f GB ⬇️\n" +
		"<b>مصرف ارسال:</b> %.2f GB ⬆️\n" +
		"<b>انقضا:</b> %s 📅</blockquote>",

	AccountRemoved: rlm + "<blockquote>لینک اکانت از چت تلگرام شما حذف شد. 🗑</blockquote>",
	NotLinked:      rlm + "<blockquote>این اکانت به چت تلگرام شما لینک نیست. ❓</blockquote>",
	UnknownCommand: rlm + "<blockquote>دستور نامعتبر. از دکمه‌های منو استفاده کنید. 🤔</blockquote>",

	LowQuotaWarning: rlm + "<blockquote><b>هشدار حجم اکانت</b> <code>%s</code>: فقط %d مگابایت باقی مانده. لطفاً برای تمدید اقدام کنید. 🔔</blockquote>",

	LanguagePicked:    rlm + "<blockquote>زبان با موفقیت تغییر کرد. ✅</blockquote>",
	SessionTimedOut:   rlm + "<blockquote>جلسه منقضی شد. لطفاً از منوی اصلی دوباره تلاش کنید. ⌛</blockquote>",
	OcservDeactivated: rlm + "<blockquote>این اکانت غیرفعال است. ⛔</blockquote>",
	RateLimited:       rlm + "<blockquote>تعداد تلاش‌ها زیاد است. لطفاً یک دقیقه صبر کنید. 🚦</blockquote>",

	// Admin
	AdminWelcome: rlm + "<blockquote><b>پنل ادمین — %s 🛡</b>\n\n" +
		"شما به‌عنوان <b>ادمین</b> وارد شده‌اید.\n" +
		"از دکمه‌های زیر برای مدیریت درخواست‌ها و نظارت بر ربات استفاده کنید.</blockquote>",
	AdminMenu: rlm + "<blockquote><b>پنل ادمین 🛡</b>\nیک گزینه انتخاب کنید:</blockquote>",

	BtnAdminPending:  rlm + "درخواست‌های در انتظار 📥",
	BtnAdminReceipts: rlm + "رسیدهای پرداخت 🧾",
	BtnAdminStats:    rlm + "آمار ربات 📊",
	BtnAdminUserView: rlm + "نمای کاربر عادی 👤",
	BtnAdminBack:     rlm + "بازگشت به پنل ادمین 🔙",
	BtnOpenPanel:     rlm + "پنل وب 🌐",

	AdminNoPending:  rlm + "<blockquote>درخواست در انتظاری وجود ندارد. 📭</blockquote>",
	AdminNoReceipts: rlm + "<blockquote>رسیدی برای تایید وجود ندارد. 📭</blockquote>",

	AdminStatsText: rlm + "<blockquote><b>آمار ربات 📊</b>\n\n" +
		"• اکانت‌های لینک‌شده: <b>%d</b>\n" +
		"• پکیج‌های فعال: <b>%d</b>\n" +
		"• درخواست‌های در انتظار: <b>%d</b>\n" +
		"• در انتظار پرداخت: <b>%d</b>\n" +
		"• رسیدهای آپلودشده: <b>%d</b></blockquote>",

	AdminRequestRow: "<b>#%d</b> · %s · <code>%s</code>\n" +
		"%s 📝\n" +
		"%s 🕒",
}
