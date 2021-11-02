export const formatDate = (date = Date.now()) => {
    const dtf = new Intl.DateTimeFormat('en', {
        timeStyle: 'short',
        hour12: false,
        dateStyle: 'long',
    })
    return dtf.format(new Date(date))
}