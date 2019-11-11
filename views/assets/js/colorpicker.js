const pickr2 = new Pickr({
    el: '#color-picker-2',
    useAsButton: true,
    default: "123456",

    components: {
        preview: true,
        opacity: true,
        hue: true,

        interaction: {
            hex: true,
            rgba: true,
            hsla: true,
            hsva: true,
            cmyk: true,
            input: true,
            clear: true,
            save: true
        }
    }
});